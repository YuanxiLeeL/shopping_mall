import { userCurrent } from "@/api/user";
import myaxios from "@/request";
import { message } from "ant-design-vue";
import { AxiosHeaderValue, AxiosResponse, AxiosResponseHeaders } from "axios";
import { defineStore } from "pinia";
import { ref } from "vue";

//远程获取用户登录信息
// export const useLoginUserStore = defineStore("loginUser", () => {
//   const loginUser = ref<any>({
//     username: "未登录",
//   });

//   async function fetchLoginUser(): Promise<void> {
//     const res = await userCurrent();
//     if (res.data.code === 0 && res.data.data) {
//       loginUser.value = res.data.data;
//     } else {
//       setTimeout(() => {
//         loginUser.value = { userName: "123", id: 1 };
//       }, 3000);
//     }
//   }

//   function setLoginUser(newLoginUser: any): void {
//     loginUser.value = newLoginUser;
//   }

//   return { loginUser, fetchLoginUser, setLoginUser };
// });

export const useLoginUserStore = defineStore("loginUser", () => {
  // 使用 ref 来定义响应式状态
  const loginUser = ref<any>({
    Username: "未登录",
    Email: "",
    PhoneNum: "",
  });

  // 获取当前用户信息
  async function fetchLoginUser(): Promise<void> {
    // 调用 userCurrent 函数获取用户信息
    const res = await userCurrent();
    if (res.data.status === "10000" && res.data.data) {
      console.log(res.data.data);
      loginUser.value = {
        ...res.data.data, // 将username存储在用户信息中
      };
    } else {
      message.error("未登录");
      setTimeout(() => {
        loginUser.value = { userName: "错误信息", id: 1 };
      }, 3000);
    }
  }

  // 手动设置用户信息
  function setLoginUser(newLoginUser: User | null): void {
    loginUser.value = newLoginUser;
  }

  return { loginUser, fetchLoginUser, setLoginUser };
});
