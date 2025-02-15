import myaxios from "@/request";
import { AxiosResponse } from "axios";

// 创建分类
export const createCategory = async (
  params: CategoryParams
): Promise<Category> => {
  try {
    const res: AxiosResponse<Category> = await myaxios.request({
      url: "/api/category",
      method: "POST",
      data: params,
    });
    return res.data;
  } catch (error) {
    console.error("创建分类失败：", error);
    throw new Error("创建分类失败，请稍后重试");
  }
};

// 获取分类列表
export const getCategories = async (): Promise<CategoryResponse> => {
  try {
    const res: AxiosResponse<CategoryResponse> = await myaxios.request({
      url: "/api/category",
      method: "GET",
    });
    return res.data;
  } catch (error) {
    console.error("获取分类列表失败：", error);
    throw new Error("获取分类列表失败，请稍后重试");
  }
};

// 获取分类下的商品
export const getGoodsByCategory = async (
  categoryname: string
): Promise<GoodsResponse> => {
  try {
    const res: AxiosResponse<GoodsResponse> = await myaxios.request({
      url: `/api/category/${categoryname}`,
      method: "GET",
    });
    return res.data;
  } catch (error) {
    console.error("获取分类下的商品失败：", error);
    throw new Error("获取分类下的商品失败，请稍后重试");
  }
};
//作为参考
// export const getGoods = async (goodName?: string): Promise<GoodsResponse> => {
//   try {
//     const res: AxiosResponse<GoodsResponse> = await myaxios.request({
//       url: "/api/goods",
//       method: "GET",
//       params: goodName ? { goodName } : {},
//     });
//     return res.data;
//   } catch (error) {
//     console.error("获取商品列表失败：", error);
//     throw new Error("获取商品列表失败，请稍后重试");
//   }
// };

// 删除分类
export const deleteCategory = async (id: string): Promise<void> => {
  try {
    await myaxios.request({
      url: `/api/category/${id}`,
      method: "DELETE",
    });
  } catch (error) {
    console.error("删除分类失败：", error);
    throw new Error("删除分类失败，请稍后重试");
  }
};
