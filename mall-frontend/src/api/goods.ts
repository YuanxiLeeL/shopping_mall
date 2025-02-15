import myaxios from "@/request";
import { AxiosResponse } from "axios";

export const getGoods = async (goodName?: string): Promise<GoodsResponse> => {
  try {
    const res: AxiosResponse<GoodsResponse> = await myaxios.request({
      url: "/api/goods",
      method: "GET",
      params: goodName ? { goodName } : {},
    });
    return res.data;
  } catch (error) {
    console.error("获取商品列表失败：", error);
    throw new Error("获取商品列表失败，请稍后重试");
  }
};

// 创建商品
export const createGood = async (params: Good): Promise<Good> => {
  try {
    const res: AxiosResponse<Good> = await myaxios.request({
      url: "/api/goods",
      method: "POST",
      data: params,
    });

    return res.data;
  } catch (error) {
    console.error("创建商品失败：", error);
    throw new Error("创建商品失败，请稍后重试");
  }
};

// 获取商品详情
export const getGoodsByID = async (id: string): Promise<Good> => {
  try {
    const res: AxiosResponse<Good> = await myaxios.request({
      url: `/api/goods/${id}`,
      method: "GET",
    });
    return res.data;
  } catch (error) {
    console.error("获取商品详情失败：", error);
    throw new Error("获取商品详情失败，请稍后重试");
  }
};

// 删除商品
export const deleteGoodsByName = async (name: string): Promise<void> => {
  try {
    await myaxios.request({
      url: `/api/goods/${name}`,
      method: "DELETE",
    });
  } catch (error) {
    console.error("删除商品失败：", error);
    throw new Error("删除商品失败，请稍后重试");
  }
};

// 更新商品
export const updateGood = async (
  id: string,
  params: Partial<Good>
): Promise<Good> => {
  try {
    const res: AxiosResponse<Good> = await myaxios.request({
      url: `/api/goods/${id}`,
      method: "PUT",
      data: params,
    });
    return res.data;
  } catch (error) {
    console.error("更新商品失败：", error);
    throw new Error("更新商品失败，请稍后重试");
  }
};
