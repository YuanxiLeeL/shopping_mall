<template>
  <div id="userRegisterView">
    <h2 class="title">用户注册</h2>
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
        :rules="[
          { required: true, message: '请输入邮箱' },
          { type: 'email', message: '请输入正确的邮箱地址' },
        ]"
      >
        <a-input
          v-model:value="formState.email"
          aria-placeholder="请输入邮箱"
        />
      </a-form-item>
      <a-form-item
        label="手机号"
        name="phonenum"
        :rules="[
          { required: true, message: '请输入手机号' },
          { pattern: /^1[3456789]\d{9}$/, message: '请输入正确的手机号' },
        ]"
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
<script lang="ts" setup>
import { userRegister } from "@/api/user";
import router from "@/router";
import { useLoginUserStore } from "@/store/useLoginUserStore";
import { message } from "ant-design-vue";
import { reactive } from "vue";
const loginUserStore = useLoginUserStore();

interface FormState {
  username: string;
  password: string;
  email: string;
  phonenum: string;
}

const formState = reactive<FormState>({
  username: "",
  password: "",
  email: "",
  phonenum: "",
});

// 表单提交事件
const handleSubmit = async (values: any) => {
  try {
    console.log(values);
    const res = await userRegister(formState);

    // 假设后端返回的数据结构如下：
    // { token: "your_jwt_token", user: { username: "admin" } }
    const { token, user } = res.data;

    // 将 Token 存储在 localStorage
    localStorage.setItem("token", token);

    // 将用户信息保存到全局状态中
    loginUserStore.setLoginUser(user); // 假设你有一个 setLoginUser 方法来保存用户信息

    // 通过 Token 获取用户信息
    await loginUserStore.fetchLoginUser();

    console.log("注册成功:", values);
    message.success("注册成功");
    router.push({
      path: "/",
      replace: true,
    });
  } catch (error) {
    console.error("注册失败", error);
  }
};

const onFinishFailed = (errorInfo: any) => {
  console.log("Failed:", errorInfo);
};
</script>
