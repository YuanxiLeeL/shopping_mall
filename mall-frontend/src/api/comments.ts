import myaxios from "@/request";
import { AxiosResponse } from "axios";

export const createComment = async (
  params: CommentParams
): Promise<Comment> => {
  try {
    const res: AxiosResponse<Comment> = await myaxios.request({
      url: `/api/goods/${params.good_id}/comments`,
      method: "POST",
      data: { content: params.content },
    });
    return res.data;
  } catch (error) {
    console.error("创建评论失败：", error);
    throw new Error("创建评论失败，请稍后重试");
  }
};

// 获取评论列表
export const getCommentsByGoodID = async (
  good_id: string
): Promise<CommentResponse> => {
  try {
    const res: AxiosResponse<CommentResponse> = await myaxios.request({
      url: `/api/goods/${good_id}/comments`,
      method: "GET",
    });
    return res.data;
  } catch (error) {
    console.error("获取评论列表失败：", error);
    throw new Error("获取评论列表失败，请稍后重试");
  }
};

// 更新评论
export const updateComment = async (
  id: string,
  params: Partial<Comment>
): Promise<Comment> => {
  try {
    const res: AxiosResponse<Comment> = await myaxios.request({
      url: `/api/comments/${id}`,
      method: "PUT",
      data: params,
    });
    return res.data;
  } catch (error) {
    console.error("更新评论失败：", error);
    throw new Error("更新评论失败，请稍后重试");
  }
};

// 删除评论
export const deleteComment = async (id: string): Promise<void> => {
  try {
    await myaxios.request({
      url: `/api/comments/${id}`,
      method: "DELETE",
    });
  } catch (error) {
    console.error("删除评论失败：", error);
    throw new Error("删除评论失败，请稍后重试");
  }
};
