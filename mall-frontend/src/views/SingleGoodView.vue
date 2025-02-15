<template>
  <div id="singleGoodView">
    <a-descriptions title="商品信息">
      <a-descriptions-item label="商品名称">{{
        good.name
      }}</a-descriptions-item>
      <a-descriptions-item label="描述">{{
        good.description
      }}</a-descriptions-item>
      <a-descriptions-item label="价格">{{ good.price }}</a-descriptions-item>
      <a-descriptions-item label="分类">{{
        good.category
      }}</a-descriptions-item>
    </a-descriptions>

    <a-divider orientation="left">评论</a-divider>

    <a-list item-layout="horizontal" :data-source="comments">
      <template #renderItem="{ item }">
        <a-list-item>
          <a-comment :author="item.username" :datetime="item.CreatedAt">
            <template #content>
              <div v-if="!item.editing">{{ item.content }}</div>
              <a-input
                v-else
                v-model:value="item.newContent"
                placeholder="请输入评论..."
                :rows="2"
              />
            </template>
            <template #actions>
              <span v-if="!item.editing">
                <a-button type="link" @click="checkCommentPermission(item)"
                  >修改</a-button
                >
                <a-button type="link" @click="deleteComment(item)"
                  >删除</a-button
                >
              </span>
              <span v-else>
                <a-button type="link" @click="saveComment(item)">提交</a-button>
                <a-button type="link" @click="cancelEdit(item)">取消</a-button>
              </span>
            </template>
          </a-comment>
        </a-list-item>
      </template>
    </a-list>

    <div class="comment-input-container">
      <a-input
        v-model:value="newComment"
        placeholder="请输入评论..."
        :rows="4"
        class="comment-input"
      />
      <a-button type="primary" @click="submitComment" class="submit-button">
        提交
      </a-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Comment, onMounted, ref } from "vue";
import axios from "axios";
import router from "@/router";
import { useRoute } from "vue-router";
import myaxios from "@/request";
import locale from "ant-design-vue/es/locale-provider";
// 假设你有一个 models 文件夹，里面有 Comment 接口定义

// 定义商品接口
interface Good {
  name: string;
  description: string;
  price: number;
  category: string;
}

// 定义一个响应式变量来存储商品数据
const good = ref<Good>({
  name: "",
  description: "",
  price: 0,
  category: "",
});

interface Comments {
  username: string;
  content: string;
  CreatedAt: string;
  newContent: string;
  editing: boolean;
  ID: number;
}

interface CommentResponse {
  comments: Comment[];
}

// 定义一个响应式变量来存储新的评论
const newComment = ref<string>("");

// 定义一个响应式变量来存储评论数据
const comments = ref<CommentResponse>();
const goodName = ref<string | undefined>(undefined);

// 保存评论
const saveComment = async (item: Comments) => {
  if (!item.newContent) {
    alert("评论内容不能为空");
    return;
  }
  try {
    const res = await myaxios.put(`/api/comments/${item.ID}`, {
      content: item.newContent,
    });

    if (res.status === 200) {
      item.content = item.newContent;
      item.editing = false;
      alert("评论修改成功！");
      goodName.value = route.query.goodName as string;
      // 重新获取评论列表
      fetchGoodInfo(goodName.value);
    } else if (res.status === 404) {
      alert("评论未找到");
    } else if (res.status === 403) {
      alert("你没有权限修改这条评论");
    } else {
      alert("修改评论失败，请稍后重试");
    }
  } catch (error: any) {
    console.error("修改评论失败：", error);
  }
};

const submitComment = async () => {
  if (!newComment.value) {
    alert("评论内容不能为空");
    return;
  }
  try {
    const res = await myaxios.post(`/api/comments`, {
      content: newComment.value, // 传递一个包含评论内容的对象
    });

    if (res.status === 200) {
      // 提交成功后清空输入框并更新评论
      newComment.value = "";
      comments.value.push(res.data.comment); // 假设返回的评论数据为新评论
      alert("评论提交成功！");
    }
  } catch (error: any) {
    console.error("提交评论失败：", error);
  }

  location.reload();
};

// 删除评论
const deleteComment = async (item: Comments) => {
  try {
    const res = await myaxios.delete(`/api/comments/${item.ID}`);

    if (res.status === 200) {
      comments.value = comments.value.filter(
        (comment) => comment.id !== item.ID
      );
      alert("评论删除成功！");
    }
  } catch (error: any) {
    console.error("删除评论失败：", error);
  }
  location.reload();
};

// 编辑评论
const editComment = (item: Comments) => {
  item.editing = true;
};

// 取消编辑
const cancelEdit = (item: Comments) => {
  item.editing = false;
  item.newContent = item.content;
};

// 定义一个函数来从后端获取商品数据
const fetchGoodInfo = async (goodname: string) => {
  try {
    const res = await myaxios.get(`/api/singlegood/${goodname}`);
    good.value = res.data.info;
    comments.value = res.data.comments;
    console.log("Comments data:", res.data.comments); // 添加日志输出
  } catch (error: any) {
    console.error("获取商品信息失败：", error);
    alert("获取商品信息失败，请稍后重试");
  }
};

const checkCommentPermission = async (item: Comments) => {
  try {
    const res = await myaxios.get(`/api/comments/${item.ID}/permission`);

    if (res.status === 200) {
      editComment(item);
      // 进入编辑状态
    } else {
      alert("你没有权限修改这条评论");
    }
  } catch (error: any) {
    console.error("验证权限失败：", error);
    alert("验证权限失败，请稍后重试");
  }
};
const route = useRoute();
// 在组件挂载时获取商品数据
onMounted(() => {
  goodName.value = route.query.goodName as string;
  fetchGoodInfo(goodName.value);
});
</script>

<style scoped>
.comment-input-container {
  position: fixed;
  bottom: 70px;
  left: 50%;
  transform: translateX(-50%);
  width: 90%;
  max-width: 600px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  background: #fff;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  z-index: 10;
}

.comment-input {
  width: 80%;
}

.submit-button {
  margin-left: 10px;
}
</style>
