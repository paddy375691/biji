from app_release.models import Env, Project, App, ReleaseConfig, ReleaseApply
from rest_framework import serializers
from django_filters import rest_framework as filters

class EnvSerializer(serializers.ModelSerializer):
    class Meta:
        model = Env
        fields = "__all__"
        read_only_fields = ('id',)

class ProjectSerializer(serializers.ModelSerializer):
    class Meta:
        model = Project
        fields = "__all__"
        read_only_fields = ('id',)

class AppSerializer(serializers.ModelSerializer):
    project = ProjectSerializer(read_only=True)

    class Meta:
        model = App
        fields = "__all__"
        read_only_fields = ('id',)

class ReleaseConfigSerializer(serializers.ModelSerializer):
    app = AppSerializer(read_only=True)
    env = EnvSerializer(read_only=True)

    class Meta:
        model = ReleaseConfig
        fields = "__all__"
        read_only_fields = ('id',)

# 定义过滤器
class ReleaseConfigFilter(filters.FilterSet):
    env = filters.CharFilter(field_name="env__id")
    app = filters.CharFilter(field_name="app__id")
    class Meta:
        model = ReleaseConfig
        fields = ("app", "env")

class ReleaseApplySerializer(serializers.ModelSerializer):
    release_config = ReleaseConfigSerializer(read_only=True)

    class Meta:
        model = ReleaseApply
        fields = "__all__"
        read_only_fields = ('id',)

# 定义过滤器
class ReleaseApplyFilter(filters.FilterSet):
    env = filters.CharFilter(field_name="release_config__env__id")
    app = filters.CharFilter(field_name="release_config__app__id")
    class Meta:
        model = ReleaseConfig
        fields = ("env", "app")