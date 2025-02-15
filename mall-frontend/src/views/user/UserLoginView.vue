<template>
  <div id="userLoginView">
    <h2 class="title">用户登录</h2>
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

      <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
        <a-button type="primary" html-type="submit">提交</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
<script lang="ts" setup>
import { userLogin } from "@/api/user";
import router from "@/router";
import { useLoginUserStore } from "@/store/useLoginUserStore";
import { message } from "ant-design-vue";
import { reactive } from "vue";
const loginUserStore = useLoginUserStore();

interface FormState {
  username: string;
  password: string;
}

const formState = reactive<FormState>({
  username: "",
  password: "",
});

// 表单提交事件
const handleSubmit = async (values: any) => {
  try {
    const res = await userLogin(formState);
    console.log("后端返回的响应:", res);
    if (!res || !res.data) {
      message.error("登录失败");
      console.error("后端返回数据无效");
      return;
    }

    const { status, msg, token } = res.data;

    // 将 Token 存储在 localStorage
    localStorage.setItem("token", token);

    // 将用户信息保存到全局状态中
    loginUserStore.setLoginUser(token); // 假设你有一个 setLoginUser 方法来保存用户信息

    // 通过 Token 获取用户信息
    await loginUserStore.fetchLoginUser();

    console.log("登录成功:", values);
    message.success("登录成功");
    router.push({
      path: "/",
      replace: true,
    });
  } catch (error) {
    console.error("登录失败", error);
  }
};

const onFinishFailed = (errorInfo: any) => {
  console.log("Failed:", errorInfo);
};
</script>
