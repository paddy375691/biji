# Generated by Django 3.2 on 2022-07-10 06:21

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ('system_config', '0002_notify'),
    ]

    operations = [
        migrations.CreateModel(
            name='App',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=30, verbose_name='应用名称')),
                ('english_name', models.CharField(max_length=30, unique=True, verbose_name='英文名称')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
            ],
            options={
                'verbose_name_plural': '应用管理',
                'db_table': 'app_release_app',
                'ordering': ('-id',),
            },
        ),
        migrations.CreateModel(
            name='Env',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=30, verbose_name='环境名称')),
                ('english_name', models.CharField(max_length=30, unique=True, verbose_name='英文名称')),
                ('note', models.TextField(blank=True, null=True, verbose_name='备注')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
            ],
            options={
                'verbose_name_plural': '环境管理',
                'db_table': 'app_release_env',
                'ordering': ('-id',),
            },
        ),
        migrations.CreateModel(
            name='Project',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=30, verbose_name='项目名称')),
                ('english_name', models.CharField(max_length=30, unique=True, verbose_name='英文名称')),
                ('note', models.TextField(blank=True, null=True, verbose_name='备注')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
            ],
            options={
                'verbose_name_plural': '项目管理',
                'db_table': 'app_release_project',
                'ordering': ('-id',),
            },
        ),
        migrations.CreateModel(
            name='ReleaseConfig',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('server_ids', models.JSONField(max_length=100, verbose_name='目标主机')),
                ('git_repo', models.CharField(max_length=100, verbose_name='Git仓库')),
                ('git_credential_id', models.IntegerField(blank=True, null=True, verbose_name='凭据ID')),
                ('note', models.TextField(blank=True, null=True, verbose_name='备注')),
                ('exclude_files', models.TextField(blank=True, null=True, verbose_name='部署排除文件')),
                ('global_variables', models.TextField(blank=True, null=True, verbose_name='自定义全局变量')),
                ('pre_checkout_script', models.TextField(blank=True, null=True, verbose_name='代码检出前执行脚本')),
                ('post_checkout_script', models.TextField(blank=True, null=True, verbose_name='代码检出后执行脚本')),
                ('dst_dir', models.CharField(max_length=100, verbose_name='服务器部署路径')),
                ('history_version_dir', models.CharField(max_length=100, verbose_name='历史版本路径')),
                ('history_version_number', models.IntegerField(default=7, verbose_name='历史版本保留数')),
                ('pre_deploy_script', models.TextField(blank=True, null=True, verbose_name='部署前执行脚本')),
                ('post_deploy_script', models.TextField(blank=True, null=True, verbose_name='部署后执行脚本')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('app', models.ForeignKey(on_delete=django.db.models.deletion.PROTECT, to='app_release.app', verbose_name='应用名称')),
                ('env', models.ForeignKey(on_delete=django.db.models.deletion.PROTECT, to='app_release.env', verbose_name='发布环境')),
                ('notify', models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.PROTECT, to='system_config.notify', verbose_name='消息通知')),
            ],
            options={
                'verbose_name_plural': '发布配置',
                'db_table': 'app_release_config',
                'ordering': ('-id',),
            },
        ),
        migrations.CreateModel(
            name='ReleaseApply',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('title', models.CharField(max_length=30, verbose_name='发布标题')),
                ('branch', models.CharField(max_length=100, verbose_name='代码分支')),
                ('server_ids', models.JSONField(verbose_name='目标主机')),
                ('version_id', models.CharField(max_length=30, unique=True, verbose_name='版本标识')),
                ('status', models.IntegerField(choices=[(1, '待发布'), (2, '发布中'), (3, '发布成功'), (4, '发布异常')], default=1, verbose_name='发布状态')),
                ('note', models.TextField(blank=True, null=True, verbose_name='备注')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('release_time', models.DateTimeField(auto_now=True, verbose_name='发布时间')),
                ('release_config', models.ForeignKey(on_delete=django.db.models.deletion.PROTECT, to='app_release.releaseconfig', verbose_name='发布配置')),
            ],
            options={
                'verbose_name_plural': '发布申请',
                'db_table': 'app_release_apply',
                'ordering': ('-id',),
            },
        ),
        migrations.AddField(
            model_name='app',
            name='project',
            field=models.ForeignKey(on_delete=django.db.models.deletion.PROTECT, to='app_release.project', verbose_name='所属项目'),
        ),
    ]