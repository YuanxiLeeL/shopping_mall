<template>
  <div id="GlobalHeader">
    <a-row :wrap="false">
      <a-col flex="100px"><div class="title">Yuanxi的商城</div></a-col>
      <a-col flex="auto"
        ><a-menu
          v-model:selectedKeys="current"
          mode="horizontal"
          :items="items"
          @click="doMenuClick"
      /></a-col>
      <a-col flex="80px">
        <div class="user-login-status">
          <div v-if="useLoginUser.loginUser">
            {{ useLoginUser.loginUser.Username }}
          </div>
          <div v-else>
            <a-button type="primary" href="api/auth/login"> 登录 </a-button>
          </div>
        </div></a-col
      >
    </a-row>
    <a-modal
      v-model:open="isCategoryModalVisible"
      title="选择分类"
      @ok="handleCategoryOk"
      @cancel="handleCategoryCancel"
    >
      <a-list item-layout="horizontal">
        <a-list-item
          v-for="category in categories"
          :key="category.id"
          @click="selectCategory(category.name)"
        >
          <a>{{ category.name }}</a>
        </a-list-item>
      </a-list>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import { h, onMounted, ref } from "vue";
import {
  MailOutlined,
  AppstoreOutlined,
  SettingOutlined,
  HomeOutlined,
  CrownOutlined,
} from "@ant-design/icons-vue";
import { MenuProps, message } from "ant-design-vue";
import { useRouter } from "vue-router";
import { useLoginUserStore } from "@/store/useLoginUserStore";
import myaxios from "@/request";

const router = useRouter();
const useLoginUser = useLoginUserStore();
//点击菜单后路由跳转事件
const doMenuClick = ({ key }: { key: string }) => {
  router.push({
    path: key,
  });
};
const current = ref<string[]>(["mail"]);

router.afterEach((to) => {
  current.value = [to.path];
});

const items = ref<MenuProps["items"]>([
  {
    key: "/",
    icon: () => h(HomeOutlined),
    label: "主页",
    title: "主页",
  },
  {
    key: "/auth/login",
    label: "用户登录",
    title: "用户登录",
  },
  {
    key: "/auth/register",
    label: "用户注册",
    title: "用户注册",
  },
  {
    key: "/cart",
    label: "购物车",
    title: "购物车",
  },
  {
    key: "/goods",
    label: "所有商品",
    title: "所有商品",
  },
  {
    key: "/user/info",
    label: "个人信息",
    title: "个人信息",
  },
  {
    key: "/categories",
    label: "分类",
    title: "分类",
    onClick: () => {
      isCategoryModalVisible.value = true;
    },
  },
]);

// 分类弹窗
const isCategoryModalVisible = ref(false);

const categories = ref([
  { id: 1, name: "电子产品", description: "各种电子产品" },
  { id: 2, name: "家居用品", description: "家居用品" },
  { id: 3, name: "服装", description: "各种服装" },
  // 添加更多分类
]);
// 从后端获取分类列表
const fetchCategories = async () => {
  try {
    const response = await myaxios.request({
      url: "api/category",
      method: "get",
    });
    categories.value = response.data; // 假设后端返回的数据是一个数组
  } catch (error) {
    message.error("获取分类列表失败");
    console.error("获取分类列表失败：", error);
  }
};

const selectCategory = (categoryname: string) => {
  isCategoryModalVisible.value = false;
  router.push({
    path: `/goods/category/${categoryname}`,
  });
};
//路由跳转
// const viewGoodDetail = (goodName: string) => {
//   if (!goodName) {
//     console.error("goodName is required");
//     return;
//   }
//   router.push({ path: "/good/info/:goodname", query: { goodName } });
// };

const handleCategoryOk = () => {
  isCategoryModalVisible.value = false;
};

const handleCategoryCancel = () => {
  isCategoryModalVisible.value = false;
};

// 在组件挂载时获取分类列表
onMounted(() => {
  fetchCategories();
});
</script>

<style scoped></style>
