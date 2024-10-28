from app_release.models import Env, Project, App, ReleaseConfig, ReleaseApply
from rest_framework import serializers

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

class ReleaseApplySerializer(serializers.ModelSerializer):
    release_config = ReleaseConfigSerializer(read_only=True)

    class Meta:
        model = ReleaseApply
        fields = "__all__"
        read_only_fields = ('id',)
