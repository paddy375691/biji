import json
import shutil

import ansible.constants as C
from ansible.executor.task_queue_manager import TaskQueueManager
from ansible.module_utils.common.collections import ImmutableDict
from ansible.inventory.manager import InventoryManager
from ansible.parsing.dataloader import DataLoader
from ansible.playbook.play import Play
from ansible.plugins.callback import CallbackBase
from ansible.vars.manager import VariableManager
from ansible import context
from ansible.executor.playbook_executor import PlaybookExecutor


"""
输出格式：
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
    },
    "task_name2": {
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
    },
}
"""

class ResultCallback(CallbackBase):
    def __int__(self, *args, **kwargs):
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
    def v2_runner_on_ok(self, result):
        print("[%s] -> %s" %(result._task.name, result._host.get_name()))
        self.__init_result_dict(result)
        self.result[result._task.name]['success'][result._host.get_name()] = result._result
    # 执行失败
    def v2_runner_on_failed(self, result, ignore_errors=False):
        print("[%s] -> %s，执行失败！" % (result._task.name, result._host.get_name()))
        self.__init_result_dict(result)
        self.result[result._task.name]['failed'][result._host.get_name()] = result._result
    # 主机不可达
    def v2_runner_on_unreachable(self, result):
        print("[%s] -> %s，主机不可达！" % (result._task.name, result._host.get_name()))
        self.__init_result_dict(result)
        self.result[result._task.name]['unreachable'][result._host.get_name()] = result._result

class AnsibleApi():
    # 定义默认参数，方便后面传入覆盖默认
    def __init__(self,
        connection = 'smart',  # 连接方式 local 本地方式，smart ssh方式
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
        inventory=None):  # 指定/etc/ansible/hosts

        # 将参数加载到ansible里
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

        self.loader = DataLoader()  # 数据解析器
        self.passwords = dict()
        self.results_callback = ResultCallback()  # 类实例化
        self.inventory = inventory if inventory else 'localhost'
        self.inventory = InventoryManager(loader=self.loader, sources=self.inventory)  # 加载资产清单文件
        self.variable_manager = VariableManager(loader=self.loader, inventory=self.inventory) # 变量管理器

    def command_run(self, hosts="localhost", task_list=None):
        tqm = TaskQueueManager(
            inventory=self.inventory,
            variable_manager=self.variable_manager,
            loader=self.loader,
            passwords=self.passwords,
            stdout_callback=self.results_callback,
        )

        # 创建一个adhoc
        play_source = dict(
            name="Ansible Play",
            hosts=hosts,
            gather_facts='no',
            tasks=task_list
        )
        play = Play().load(play_source, variable_manager=self.variable_manager, loader=self.loader)

        try:
            result = tqm.run(play)
        finally:
            tqm.cleanup()
            if self.loader:
                self.loader.cleanup_all_tmp_files()
        shutil.rmtree(C.DEFAULT_LOCAL_TMP, True)

    def playbook_run(self, playbook_path):
        playbook = PlaybookExecutor(
            playbooks=playbook_path,
            inventory=self.inventory,
            variable_manager=self.variable_manager,
            loader=self.loader,
            passwords=self.passwords
        )
        # 使用回调
        playbook._tqm._stdout_callback = self.results_callback
        result = playbook.run()  # 运行
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