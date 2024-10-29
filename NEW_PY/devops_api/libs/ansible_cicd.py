from ansible.plugins.callback import CallbackBase
from ansible.parsing.dataloader import DataLoader  # 用于读取YAML和JSON格式的文件
from ansible.vars.manager import VariableManager  # 用于存储各类变量信息
from ansible.inventory.manager import InventoryManager  # 用于导入资产文件
from ansible.playbook.play import Play  # 存储执行hosts的角色信息
from ansible.executor.task_queue_manager import TaskQueueManager  # ansible底层用到的任务队列
from ansible.module_utils.common.collections import ImmutableDict
from ansible import context
from ansible.executor.playbook_executor import PlaybookExecutor  # 核心类执行playbook
import ansible.constants as C
import shutil, json

# 由于官方示例值的回调函数没有进行格式化，需要重写回调插件
class ResultCallback(CallbackBase):
    """
    return data
    {
        "task_name1": {
            "success": {
                "192.168.1.71": {},
                "192.168.1.72": {}
            },
            "failed": {
                "192.168.1.71": {}
            },
            "unreachable": {
                "192.168.1.73": {}
            }
        }
    }
    """

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.result = {}

    # 初始化数据格式
    def __init_result_dict(self, result):
        self.result[result._task.name] = {
            "success": {},
            "failed": {},
            "unreachable": {}
        }

    # 执行成功
    def v2_runner_on_ok(self, result, *args, **kwargs):
        print("[%s] -> %s" % (result._task.name, result._host.get_name()))
        self.__init_result_dict(result)
        self.result[result._task.name]["success"][result._host.get_name()] = result._result

    # 执行失败
    def v2_runner_on_failed(self, result, *args, **kwargs):
        print("[%s] -> %s，执行失败！" % (result._task.name, result._host.get_name()))
        self.__init_result_dict(result)
        self.result[result._task.name]["failed"][result._host.get_name()] = result._result

    # 主机不可达
    def v2_runner_on_unreachable(self, result):
        print("[%s] -> %s，主机不可达！" % (result._task.name, result._host.get_name()))
        self.__init_result_dict(result)
        self.result[result._task.name]["unreachable"][result._host.get_name()] = result._result

class AnsibleApi():
    """
    在初始化的时候可以传参，以便覆盖默认选项的值
    """
    def __init__(self,
                 connection='smart',  # 连接方式 local 本地方式，smart ssh方式
                 remote_user=None,  # 远程用户
                 ack_pass=None,  # 提示输入密码
                 sudo=None, sudo_user=None, ask_sudo_pass=None,
                 module_path=None,  # 模块路径，可以指定一个自定义模块的路径
                 become=None,  # 是否提权
                 become_method=None,  # 提权方式 默认 sudo 可以是 su
                 become_user=None,  # 提权后，要成为的用户，并非登录用户
                 check=False, diff=False,
                 listhosts=None, listtasks=None, listtags=None,
                 verbosity=3,
                 syntax=None,
                 start_at_task=None,
                 inventory=None):
        context.CLIARGS = ImmutableDict(
            connection=connection,
            remote_user=remote_user,
            ack_pass=ack_pass,
            sudo=sudo,
            sudo_user=sudo_user,
            ask_sudo_pass=ask_sudo_pass,
            module_path=module_path,
            become=become,
            become_method=become_method,
            become_user=become_user,
            verbosity=verbosity,
            listhosts=listhosts,
            listtasks=listtasks,
            listtags=listtags,
            syntax=syntax,
            start_at_task=start_at_task,
        )

        self.loader = DataLoader()  # 数据解析器，用于解析资产清单文件（hosts）中的数据和变量。
        self.passwords = dict()  # 密码这里是必须使用的一个参数，假如通过了公钥信任，也可以给一个空字典
        self.results_callback = ResultCallback()  # 实例化回调类
        self.inventory = inventory if inventory else "localhost"
        self.inventory = InventoryManager(loader=self.loader, sources=self.inventory)  # 指定资产清单文件
        self.variable_manager = VariableManager(loader=self.loader, inventory=self.inventory)  # 变量管理器，会资产清单文件汇总获取定义好的变量

    def command_run(self, hosts='localhost', task_list=None):
        # 创建一个adhoc
        play_source = dict(
            name="Ansible Play",
            hosts=hosts,  # 指定分组
            gather_facts='no',
            tasks=task_list  # 值是列表[]，每个元素字典{}就是一个task
        )
        # 创建一个play对象
        play = Play().load(play_source, variable_manager=self.variable_manager, loader=self.loader)

        # 任务队列管理器：要想执行 Ad-hoc ，需要把上面的 play 对象交个任务队列管理器的 run 方法去运行。
        tqm = None
        try:
            tqm = TaskQueueManager(
                inventory=self.inventory,
                variable_manager=self.variable_manager,
                loader=self.loader,
                passwords=self.passwords,
                stdout_callback=self.results_callback,  # 这里是使用了之前，自定义的回调插件，而不是默认的回调插件 `default`
            )
            result = tqm.run(play)  # 执行的结果返回码，成功是 0
        finally:
            if tqm is not None:
                tqm.cleanup()

    def playbook_run(self, playbook_path):
        playbook = PlaybookExecutor(playbooks=playbook_path,
                                    inventory=self.inventory,
                                    variable_manager=self.variable_manager,
                                    loader=self.loader,
                                    passwords=self.passwords)
        # 使用回调函数，json输出
        playbook._tqm._stdout_callback = self.results_callback

        result = playbook.run()
        return result

    def get_result(self):
        return self.results_callback.result

if __name__ == "__main__":
    ansible = AnsibleApi(inventory="/etc/ansible/hosts")
    # task_list= [
    #             dict(action=dict(module='shell', args='ls'), register='shell_out'),
    #             dict(action=dict(module='debug', args=dict(msg='{{shell_out.stdout}}'))),
    #             dict(action=dict(module='command', args=dict(cmd='/usr/bin/uptime'))),
    #         ]
    # ansible.command_run(hosts="webservers", task_list=task_list)
    ansible.playbook_run(['/root/deploy.yaml'])
    print(ansible.get_result())