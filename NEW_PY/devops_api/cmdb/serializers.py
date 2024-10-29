from cmdb.models import Idc,ServerGroup,Server
from rest_framework import serializers

class IdcSerializer(serializers.ModelSerializer):
    """
    IDC机房序列化类
    """
    class Meta:
        model = Idc
        fields = "__all__"
        read_only_fields = ('id',)

class ServerGroupSerializer(serializers.ModelSerializer):
    """
    主机分组序列化类
    """
    class Meta:
        model = ServerGroup
        fields = "__all__"
        read_only_fields = ('id',)

class ServerSerializer(serializers.ModelSerializer):
    """
    服务器序列化类
    """
    idc = IdcSerializer(read_only=True)  # 一对多
    server_group = ServerGroupSerializer(many=True, read_only=True) # 多对多

    class Meta:
        model = Server
        fields = "__all__"
        read_only_fields = ('id',)
