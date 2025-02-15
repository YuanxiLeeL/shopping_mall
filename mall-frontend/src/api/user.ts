import myaxios from "@/request";
import { message } from "ant-design-vue";
import { Axios, AxiosResponse, AxiosError } from "axios";

//用户注册
export const userRegister = async (credentials: {
  username: string;
  password: string;
  email: string;
  phonenum: string;
}): Promise<AxiosResponse> => {
  const res: AxiosResponse<any, any> = await myaxios.request({
    url: "/api/auth/register",
    method: "POST",
    data: credentials,
  });
  return res;
};

//用户注册
export const userInfoEdit = async (credentials: {
  username: string;
  password: string;
  email: string;
  phonenum: string;
}): Promise<AxiosResponse> => {
  const res: AxiosResponse<any, any> = await myaxios.request({
    url: "/api/user/info/edit",
    method: "PUT",
    data: credentials,
  });
  return res;
};
//用户登录
export const userLogin = async (credentials: {
  username: string;
  password: string;
}): Promise<any> => {
  try {
    const res = await myaxios.post("/api/auth/login", credentials);
    if (res.status === 401) {
      message.error("用户名或密码错误");
      return;
    }
    return res; // 返回后端返回的数据
  } catch (error: any) {
    console.error("登录失败", error);
    throw error; // 重新抛出错误
  }
};
//用户注销
// export const userLogout = async (): Promise<AxiosResponse> => {
//   const res: AxiosResponse<any, any> = await myaxios.request({
//     url: "/api/user/logout",
//     method: "POST",
//   });
//   return res;
// };

export const userCurrent = async () => {
  return myaxios.request({
    url: "/api/auth/current",
    method: "GET",
  });
};

// 更新用户名
export const updateUsername = async (username: string): Promise<void> => {
  try {
    await myaxios.request({
      url: "/api/auth/username",
      method: "PUT",
      data: { username },
    });
  } catch (error) {
    console.error("更新用户名失败：", error);
    throw new Error("更新用户名失败，请稍后重试");
  }
};

// 验证密码
export const authPassword = async (password: string): Promise<void> => {
  try {
    await myaxios.request({
      url: "/api/auth/authpassword",
      method: "POST",
      data: { password },
    });
  } catch (error) {
    console.error("验证密码失败：", error);
    throw new Error("验证密码失败，请稍后重试");
  }
};

// 更新密码
export const updatePassword = async (password: string): Promise<void> => {
  try {
    await myaxios.request({
      url: "/api/auth/password",
      method: "PUT",
      data: { password },
    });
  } catch (error) {
    console.error("更新密码失败：", error);
    throw new Error("更新密码失败，请稍后重试");
  }
};
