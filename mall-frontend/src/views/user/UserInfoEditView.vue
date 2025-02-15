<template>
  <div id="userInfoEditView">
    <h2 class="title">信息修改</h2>
    <a-form
      :model="formState"
      name="basic"
      :label-col="{ span: 8 }"
      :wrapper-col="{ span: 16 }"
      autocomplete="off"
      @finish="handleSubmit"
      @finishFailed="onFinishFailed"
    >
      <a-form-item
        label="账号"
        name="username"
        :rules="[{ required: true, message: '请输入账号' }]"
      >
        <a-input
          v-model:value="formState.username"
          aria-placeholder="请输入账号"
        />
      </a-form-item>

      <a-form-item
        label="密码"
        name="password"
        :rules="[
          { required: true, message: '请输入密码' },
          { min: 3, message: '密码长度至少为3' },
        ]"
      >
        <a-input-password
          v-model:value="formState.password"
          aria-placeholder="请输入密码"
        />
      </a-form-item>
      <a-form-item
        label="邮箱"
        name="email"
        :rules="[{ required: true, message: '请输入邮箱' }]"
      >
        <a-input
          v-model:value="formState.email"
          aria-placeholder="请输入邮箱"
        />
      </a-form-item>
      <a-form-item
        label="手机号"
        name="phonenum"
        :rules="[{ required: true, message: '请输入手机号' }]"
      >
        <a-input
          v-model:value="formState.phonenum"
          aria-placeholder="请输入手机号"
        />
      </a-form-item>

      <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
        <a-button type="primary" html-type="submit">提交</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script>
import { defineComponent, reactive } from "vue";
import { message } from "ant-design-vue";
import { userInfoEdit } from "@/api/user";
import router from "@/router";
export default defineComponent({
  setup() {
    const formState = reactive({
      username: "", // 账号
      password: "", // 密码
      email: "", // 邮箱
      phonenum: "", // 手机号
    });

    // 提交表单
    const handleSubmit = async (values) => {
      console.log("提交的数据：", values);
      const res = await userInfoEdit(formState);
      // 这里可以发送请求到后端
      message.success("提交成功");
      router.push({ name: "登录" });
      message.info("请重新登录");
    };

    // 表单验证失败
    const onFinishFailed = (errorInfo) => {
      console.log("表单验证失败：", errorInfo);
      message.error("表单验证失败");
    };

    return {
      formState,
      handleSubmit,
      onFinishFailed,
    };
  },
});
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  background-color: #f9f9f9;
}

.title {
  text-align: center;
  margin-bottom: 20px;
}

.ant-form-item {
  margin-bottom: 15px;
}

.ant-btn-primary {
  width: 100%;
}
</style>
