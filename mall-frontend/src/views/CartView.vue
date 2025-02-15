<template>
  <div id="cartView">
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
            <a @click="viewGoodDetail(record.goodname)">查看商品详情页</a>
            <a-divider type="vertical" />
            <a @click="doDelete(record.goodname)">删除</a>
          </span>
        </template>
      </template>
    </a-table>
    <a-button type="primary" @click="placeOrder">结算</a-button>
  </div>
</template>
<script lang="ts" setup>
import { getCart, removeFromCart } from "@/api/cart";
import { deleteGoodsByName, getGoods } from "@/api/goods";
import myaxios from "@/request";
import router from "@/router";
import { SmileOutlined, DownOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import axios from "axios";
import { ref } from "vue";
import { useRoute } from "vue-router";

const searchValue = ref("");
//删除商品
const doDelete = async (name: string) => {
  if (!name) {
    return;
  }
  const res = await removeFromCart(name);
};

//路由跳转
const viewGoodDetail = (goodName: string) => {
  if (!goodName) {
    console.error("goodName is required");
    return;
  }
  router.push({ path: "/good/info/:goodname", query: { goodName } });
};

const route = useRoute();

const placeOrder = async () => {
  const res = await myaxios.get("api/placeorder");
  if (res) {
    message.success("下单成功,总金额为" + res.data.totalAmount);
  } else {
    message.error("下单失败");
  }
};

const columns = [
  {
    title: "商品名",
    dataIndex: "goodname",
  },
  {
    title: "数量",
    dataIndex: "quantity",
  },
  {
    title: "价格",
    dataIndex: "price",
  },

  //   {
  //     title: "描述",
  //     dataIndex: "description",
  //   },
  //   {
  //     title: "分类",
  //     dataIndex: "category",
  //   },
  {
    title: "操作",
    key: "action",
  },
];

const data = ref();

const fetchCart = async () => {
  //   axios.get("http://localhost:8080/api/cart").then((res) => {
  //     console.log(res, "测试");
  //   });
  const resData = await getCart();
  console.log(Array.isArray(resData));
  console.log(resData);
  if (resData) {
    data.value = resData;
  } else {
    message.error("获取商品列表失败");
  }
};

fetchCart();
</script>
