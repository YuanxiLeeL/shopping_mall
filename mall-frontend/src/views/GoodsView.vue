<template>
  <div id="goodsView">
    <a-input-search
      v-model:value="searchValue"
      placeholder="输入商品名搜索"
      enter-button="搜索"
      size="large"
      @search="onSearch"
    />
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
<script lang="ts" setup>
import { addToCart } from "@/api/cart";
import { deleteGoodsByName, getGoods } from "@/api/goods";
import router from "@/router";
import { SmileOutlined, DownOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import axios from "axios";
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";

const searchValue = ref("");
const onSearch = () => {
  // 搜索
  fetchGoods(searchValue.value);
};

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

//删除商品
const doDelete = async (name: string) => {
  if (!name) {
    return;
  }
  const res = await deleteGoodsByName(name);
};
//路由跳转
const viewGoodDetail = (goodName: string) => {
  if (!goodName) {
    console.error("goodName is required");
    return;
  }
  router.push({ path: "/good/info/:goodname", query: { goodName } });
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
    title: "分类",
    dataIndex: "category",
  },
  {
    title: "操作",
    key: "action",
  },
];

const data = ref();

const fetchGoods = async (goodName?: string) => {
  const resData = await getGoods(goodName);
  console.log(Array.isArray(resData));
  console.log(resData);
  if (resData) {
    data.value = resData;
  } else {
    message.error("获取商品列表失败");
  }
};

const good = ref();
const comments = ref([]);
// const fetchGoodInfo = async (goodName: string) => {
//   try {
//     const res = await axios.get(`/api/goods/${goodName}`);
//     good.value = res.data.info;
//     comments.value = res.data.comments;
//   } catch (error) {
//     console.error("获取商品信息失败：", error);
//   }
// };
const route = useRoute();
// onMounted(() => {
//   const goodName = route.params.goodName as string;
//   goodName;
// });
fetchGoods();
</script>
