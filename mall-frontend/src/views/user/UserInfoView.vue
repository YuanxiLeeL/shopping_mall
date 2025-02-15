<template>
  <div id="userInfoView">
    <a-descriptions title="用户信息">
      <a-descriptions-item label="用户名">{{
        userData.Username
      }}</a-descriptions-item>
      <a-descriptions-item label="邮箱地址">{{
        userData.Email
      }}</a-descriptions-item>
      <a-descriptions-item label="手机号">{{
        userData.PhoneNum
      }}</a-descriptions-item>
    </a-descriptions>
    <a-button type="primary" @click="showPasswordModal = true">
      修改用户信息
    </a-button>
    <a-modal
      :open="showPasswordModal"
      title="输入密码进行验证"
      @ok="verifyPassword"
      @cancel="showPasswordModal = false"
    >
      <a-form-item
        label="密码"
        name="password"
        :rules="[{ message: '请输入密码' }]"
      >
        <a-input-password v-model:value="password" placeholder="请输入密码" />
      </a-form-item>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import { FormInstance } from "ant-design-vue";
import { message } from "ant-design-vue";
import myaxios from "@/request";
import { useRoute } from "vue-router";
import router from "@/router";
const userData = ref({
  Username: "",
  Email: "",
  PhoneNum: "",
});

// 显示密码输入弹窗的标志
const showPasswordModal = ref(false);
// 存储用户输入的密码
const password = ref("");

const verifyPassword = async () => {
  try {
    const response = await myaxios.request({
      url: "/api/user/authpassword",
      method: "post",
      data: { old_password: password.value }, // 发送符合后端要求的 JSON 对象
    });
    if (response.data.message) {
      // 密码验证通过，跳转到修改页面
      router.push({ name: "修改用户信息" });
      showPasswordModal.value = false;
    } else {
      message.error("密码验证失败，请重新输入");
    }
  } catch (error) {
    message.error("密码不正确");
    console.error("密码验证出错:", error);
  }
};
// 获取用户信息
const fetchUserInfo = async () => {
  try {
    const response = await myaxios.get("/api/auth/current");
    userData.value = response.data.data;
  } catch (error) {
    message.error("获取用户信息失败");
  }
};

// 提交修改
const submitForm = async () => {
  try {
    const response = await axios.put("/api/user/update", userData.value);
    message.success("用户信息修改成功");
  } catch (error) {
    message.error("用户信息修改失败");
  }
};

onMounted(() => {
  fetchUserInfo();
});
</script>

<style scoped>
/* 可添加自定义样式 */
</style>
