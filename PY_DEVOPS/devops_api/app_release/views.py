from rest_framework.viewsets import ModelViewSet
from app_release.models import Env, Project, App, ReleaseConfig, ReleaseApply
from system_config.models import Notify
from app_release.serializers import EnvSerializer, ProjectSerializer, AppSerializer, ReleaseConfigSerializer,ReleaseApplySerializer

from rest_framework import filters
from django_filters.rest_framework import DjangoFilterBackend
from rest_framework.views import APIView

from rest_framework.response import Response
from django.conf import settings

class EnvViewSet(ModelViewSet):
    queryset = Env.objects.all()  # 指定操作的数据
    serializer_class = EnvSerializer  # 指定序列化器

    filter_backends = [filters.SearchFilter, filters.OrderingFilter, DjangoFilterBackend]  # 指定过滤器
    search_fields = ('name','english_name')  # 指定可搜索的字段

    def create(self, request, *args, **kwargs):
        serializer = self.get_serializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        self.perform_create(serializer)
        res = {'code': 200, 'msg': '创建成功'}
        return Response(res)

    def update(self, request, *args, **kwargs):
        partial = kwargs.pop('partial', False)
        instance = self.get_object()
        serializer = self.get_serializer(instance, data=request.data, partial=partial)
        serializer.is_valid(raise_exception=True)
        self.perform_update(serializer)
        res = {'code': 200, 'msg': '更新成功'}
        return Response(res)

    def destroy(self, request, *args, **kwargs):
        instance = self.get_object()
        try:
            self.perform_destroy(instance)
            res = {'code': 200, 'msg': '删除成功'}
            return Response(res)
        except Exception as e:
            res = {'code': 500, 'msg': '该环境关联其他发布配置，请先删除关联的发布配置再操作！'}
            return Response(res)

class ProjectViewSet(ModelViewSet):
    queryset = Project.objects.all()
    serializer_class = ProjectSerializer

    filter_backends = [filters.SearchFilter, filters.OrderingFilter, DjangoFilterBackend]
    search_fields = ('name','english_name')

    def create(self, request, *args, **kwargs):
        serializer = self.get_serializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        self.perform_create(serializer)
        res = {'code': 200, 'msg': '创建成功'}
        return Response(res)

    def update(self, request, *args, **kwargs):
        partial = kwargs.pop('partial', False)
        instance = self.get_object()
        serializer = self.get_serializer(instance, data=request.data, partial=partial)
        serializer.is_valid(raise_exception=True)
        self.perform_update(serializer)
        res = {'code': 200, 'msg': '更新成功'}
        return Response(res)

    def destroy(self, request, *args, **kwargs):
        instance = self.get_object()
        try:
            self.perform_destroy(instance)
            res = {'code': 200, 'msg': '删除成功'}
            return Response(res)
        except Exception as e:
            res = {'code': 500, 'msg': '该项目关联其他应用，请先删除关联的应用再操作！'}
            return Response(res)

class AppViewSet(ModelViewSet):
    queryset = App.objects.all()
    serializer_class = AppSerializer

    filter_backends = [filters.SearchFilter, filters.OrderingFilter, DjangoFilterBackend]  # 指定过滤器
    search_fields = ('name','english_name')

    def create(self, request, *args, **kwargs):
        serializer = self.get_serializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        # self.perform_create(serializer)

        # 一对多
        project = request.data.get('project')
        project_obj = Project.objects.get(id=project)
        App.objects.create(
            name=request.data.get('name'),
            english_name=request.data.get('english_name'),
            project=project_obj
        )

        res = {'code': 200, 'msg': '创建成功'}
        return Response(res)

    def update(self, request, pk=None, *args, **kwargs):
        partial = kwargs.pop('partial', False)
        instance = self.get_object()
        serializer = self.get_serializer(instance, data=request.data, partial=partial)
        serializer.is_valid(raise_exception=True)
        # self.perform_update(serializer)

        # 处理一对多关系
        project = request.data.get('project')
        project_obj = Project.objects.get(id=project)
        del request.data['project']
        App.objects.filter(
            id=pk
        ).update(
            project=project_obj,
            **request.data
        )

        res = {'code': 200, 'msg': '更新成功'}
        return Response(res)

    def destroy(self, request, *args, **kwargs):
        instance = self.get_object()
        self.perform_destroy(instance)
        res = {'code': 200, 'msg': '删除成功'}
        return Response(res)

class ReleaseConfigViewSet(ModelViewSet):
    queryset = ReleaseConfig.objects.all()
    serializer_class = ReleaseConfigSerializer

    filter_backends = [filters.SearchFilter, filters.OrderingFilter, DjangoFilterBackend]  # 指定过滤器
    search_fields = ('name','english_name')

    def create(self, request, *args, **kwargs):
        serializer = self.get_serializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        # self.perform_create(serializer)

        # 一对多
        app = request.data.get('app')
        env = request.data.get('env')
        notify = request.data.get('notify')
        app_obj = App.objects.get(id=app)
        env_obj = Env.objects.get(id=env)
        notify_obj = Notify.objects.get(id=notify)
        del request.data['app']
        del request.data['env']
        del request.data['notify']
        ReleaseConfig.objects.create(
            app=app_obj,
            env=env_obj,
            notify=notify_obj,
            **request.data
        )
        res = {'code': 200, 'msg': '创建成功'}
        return Response(res)

    def update(self, request, pk=None, *args, **kwargs):
        partial = kwargs.pop('partial', False)
        instance = self.get_object()
        serializer = self.get_serializer(instance, data=request.data, partial=partial)
        serializer.is_valid(raise_exception=True)
        # self.perform_update(serializer)

        app = request.data.get('app')
        env = request.data.get('env')
        notify = request.data.get('notify')
        app_obj = App.objects.get(id=app)
        env_obj = Env.objects.get(id=env)
        notify_obj = Notify.objects.get(id=notify)
        del request.data['app']
        del request.data['env']
        del request.data['notify']
        ReleaseConfig.objects.filter(
            id=pk
        ).update(
            app=app_obj,
            env=env_obj,
            notify=notify_obj,
            **request.data
        )

        res = {'code': 200, 'msg': '更新成功'}
        return Response(res)

    def destroy(self, request, *args, **kwargs):
        instance = self.get_object()
        try:
            self.perform_destroy(instance)
            res = {'code': 200, 'msg': '删除成功'}
            return Response(res)
        except Exception as e:
            res = {'code': 500, 'msg': '该发布配置关联其他发布申请，请先删除关联的发布申请再操作！'}
            return Response(res)

class ReleaseApplyViewSet(ModelViewSet):
    queryset = ReleaseApply.objects.all()
    serializer_class = ReleaseApplySerializer

    filter_backends = [filters.SearchFilter, filters.OrderingFilter, DjangoFilterBackend]
    search_fields = ('name','english_name')

    def create(self, request, *args, **kwargs):
        serializer = self.get_serializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        # self.perform_create(serializer)

        release_config = request.data.get('release_config')
        config_obj = ReleaseConfig.objects.get(id=release_config)
        del request.data['release_config']
        ReleaseApply.objects.create(
            release_config=config_obj,
            **request.data
        )

        res = {'code': 200, 'msg': '创建成功'}
        return Response(res)

    # def update(self, request, pk=None, *args, **kwargs):
    #     partial = kwargs.pop('partial', False)
    #     instance = self.get_object()
    #     serializer = self.get_serializer(instance, data=request.data, partial=partial)
    #     serializer.is_valid(raise_exception=True)
    #     # self.perform_update(serializer)
    #
    #     release_config = request.data.get('release_config')
    #     config_obj = ReleaseConfig.objects.get(id=release_config)
    #     del request.data['release_config']
    #     ReleaseApply.objects.filter(
    #         id=pk
    #     ).update(
    #         release_config=config_obj,
    #         **request.data
    #     )
    #     res = {'code': 200, 'msg': '更新成功'}
    #     return Response(res)

    def destroy(self, request, *args, **kwargs):
        instance = self.get_object()
        self.perform_destroy(instance)
        res = {'code': 200, 'msg': '删除成功'}
        return Response(res)

class GitView(APIView):
    def get(self, request):
        git_repo = request.query_params.get('git_repo')

        from git.repo import Repo
        import os, shutil

        clone_dir = os.path.join(settings.BASE_DIR, 'repo')

        if os.path.exists(clone_dir):
            shutil.rmtree(clone_dir)

        Repo.clone_from(git_repo, to_path=clone_dir)
        # 获取分支
        repo = Repo(clone_dir)
        branch = []
        for ref in repo.remote().refs:
            if ref.remote_head != 'HEAD':
                branch.append(ref.remote_head)

        res = {'code': 200, 'msg': '获取成功', 'data': branch}
        return Response(res)
