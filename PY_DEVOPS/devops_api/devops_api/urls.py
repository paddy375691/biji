"""devops_api URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/3.2/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include, re_path

from libs.token_auth import CustomAuthToken,ChangeUserPasswordView
from cmdb.views import HostCollectView, CreateHostView, ExcelCreateHostView, AliyunCloudView, TencentCloudView
from app_release.views import GitView

from rest_framework_swagger.views import get_swagger_view
schema_view = get_swagger_view(title='接口文档')

urlpatterns = [
    path('admin/', admin.site.urls),
    re_path('^api/docs/$', schema_view),
    re_path('^api/login/$', CustomAuthToken.as_view()),
    re_path('^api/change_password/$', ChangeUserPasswordView.as_view()),
    re_path('^api/cmdb/host_collect/$', HostCollectView.as_view()),
    re_path('^api/cmdb/create_host/$', CreateHostView.as_view()),
    re_path('^api/cmdb/excel_create_host/$', ExcelCreateHostView.as_view()),
    re_path('^api/cmdb/aliyun_cloud/$', AliyunCloudView.as_view()),
    re_path('^api/cmdb/tencent_cloud/$', TencentCloudView.as_view()),
    re_path('^api/app_release/git/$', GitView.as_view()),
]

# CMDB系统
from cmdb.views import IdcViewSet, ServerViewSet, ServerGroupViewSet
from rest_framework import routers
router = routers.DefaultRouter()
router.register(r'cmdb/idc', IdcViewSet, basename='idc')
router.register(r'cmdb/server_group', ServerGroupViewSet, basename='server_group')
router.register(r'cmdb/server', ServerViewSet, basename='server')

# 发布系统
from app_release.views import EnvViewSet, ProjectViewSet, AppViewSet, ReleaseConfigViewSet, ReleaseApplyViewSet
router.register(r'app_release/env', EnvViewSet, basename='env')
router.register(r'app_release/project', ProjectViewSet, basename='project')
router.register(r'app_release/app', AppViewSet, basename='app')
router.register(r'app_release/config', ReleaseConfigViewSet, basename='config')
router.register(r'app_release/apply', ReleaseApplyViewSet, basename='apply')

# 系统配置
from system_config.views import CredentialViewSet, NotifyViewSet
router.register(r'config/credential', CredentialViewSet, basename='credential')
router.register(r'config/notify', NotifyViewSet, basename='notify')

urlpatterns += [
    path('api/', include(router.urls))
]

