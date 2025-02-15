<template>
  <div id="categoryView">
    <a-table :columns="columns" :data-source="data">
      <template #headerCell="{ column }">
        <template v-if="column.key === 'name'">
          <span>
            <smile-outlined />
            Name
          </span>
        </template>
      </template>

      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'name'">
          <a>
            {{ record.name }}
          </a>
        </template>
        <template v-else-if="column.key === 'tags'">
          <span>
            <a-tag
              v-for="tag in record.tags"
              :key="tag"
              :color="
                tag === 'loser'
                  ? 'volcano'
                  : tag.length > 5
                  ? 'geekblue'
                  : 'green'
              "
            >
              {{ tag.toUpperCase() }}
            </a-tag>
          </span>
        </template>
        <template v-else-if="column.key === 'action'">
          <span>
            <a @click="viewGoodDetail(record.name)">查看商品详情页面</a>
            <a-divider type="vertical" />
            <a @click="doAdd(record.name)">加入购物车</a>
          </span>
        </template>
      </template>
    </a-table>
  </div>
</template>

<script setup lang="ts">
import { addToCart } from "@/api/cart";
import { getGoodsByCategory } from "@/api/category";
import router from "@/router";
import { message } from "ant-design-vue";
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";

const doAdd = async (name: string) => {
  if (!name) {
    return;
  }
  const res = await addToCart(name);
  if (res) {
    message.success("添加成功");
  } else {
    message.error("添加失败");
  }
};
const columns = [
  {
    title: "商品名",
    dataIndex: "name",
  },
  {
    title: "价格",
    dataIndex: "price",
  },
  {
    title: "描述",
    dataIndex: "description",
  },
  {
    title: "操作",
    key: "action",
  },
];

// 获取路由参数
const route = useRoute();
const data = ref();
const category = ref<string | undefined>(route.params.categoryname as string);

// 定义一个函数来从后端获取商品数据
// const fetchGoodInfo = async (goodname: string) => {
//   try {
//     const res = await myaxios.get(`/api/singlegood/${goodname}`);
//     good.value = res.data.info;
//     comments.value = res.data.comments;
//     console.log("Comments data:", res.data.comments); // 添加日志输出
//   } catch (error: any) {
//     console.error("获取商品信息失败：", error);
//     alert("获取商品信息失败，请稍后重试");
//   }
// };
const fetchGoodsByCategory = async (category: string) => {
  console.log(category);
  const resData = await getGoodsByCategory(category);
  console.log(Array.isArray(resData));
  console.log(resData);
  if (resData) {
    data.value = resData;
  } else {
    message.error("获取商品列表失败");
  }
};

//路由跳转
const viewGoodDetail = (goodName: string) => {
  if (!goodName) {
    console.error("goodName is required");
    return;
  }
  router.push({ path: "/good/info/:goodname", query: { goodName } });
};

// 根据 URL 中的 category 参数获取商品列表
onMounted(() => {
  console.log(category.value);
  if (category.value) {
    fetchGoodsByCategory(category.value);
  } else {
    message.error("未指定分类");
    console.log(category.value);
  }
});
</script>
