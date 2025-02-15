import myaxios from "@/request";
import { AxiosResponse } from "axios";

// 添加到购物车
export const addToCart = async (name: string) => {
  try {
    const res: AxiosResponse<CartItem> = await myaxios.request({
      url: `/api/cart/${name}`,
      method: "POST",
    });
    return res.data;
  } catch (error) {
    console.error("添加到购物车失败：", error);
    throw new Error("添加到购物车失败，请稍后重试");
  }
};

// 从购物车移除
export const removeFromCart = async (goodname: string) => {
  try {
    const res: AxiosResponse<CartResponse> = await myaxios.request({
      url: `/api/cart/${goodname}`,
      method: "DELETE",
    });
    return res;
  } catch (error) {
    console.error("从购物车移除失败：", error);
    throw new Error("从购物车移除失败，请稍后重试");
  }
}; //

// 获取购物车
export const getCart = async (): Promise<CartResponse> => {
  try {
    const res: AxiosResponse<CartResponse> = await myaxios.request({
      url: "/api/cart",
      method: "GET",
    });
    // console.log(res.data);
    return res.data;
  } catch (error) {
    console.error("获取购物车失败：", error);
    throw new Error("获取购物车失败，请稍后重试");
  }
};
