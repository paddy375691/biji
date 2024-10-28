<template>
    这是首页
</template>

<script>
    export default {
        name: "Dashboard",
        data() {
            return {
                dialogFormVisible: false, // 显示与隐藏
                form: {
                    name: '',
                    age: '',
                    sex: ''
                },
            }
        },
        methods: {
            open() {
                this.$message.success('登录成功')
            },

            msgBox() {
                this.$messageBox.confirm(
                  '你确认要删除这个吗？',
                  '提示',
                  {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning',
                  }
                )
                  .then(() => {
                    this.$message({
                      type: 'success',
                      message: '删除成功',
                    })
                  })
                  .catch(() => {
                    this.$message({
                      type: 'info',
                      message: '删除取消',
                    })
                  })
              }

        }
    }
</script>

<style scoped>

</style>