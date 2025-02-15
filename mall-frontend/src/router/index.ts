import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import HomeView from "../views/HomeView.vue";
import UserLoginView from "@/views/user/UserLoginView.vue";
import { userRegister } from "@/api/user";
import UserRegisterView from "@/views/user/UserRegisterView.vue";

import GoodsView from "@/views/GoodsView.vue";
import CartView from "@/views/CartView.vue";
import GoodsManageView from "@/views/GoodsManageView.vue";
import SingleGoodView from "@/views/SingleGoodView.vue";
import UserInfoView from "@/views/user/UserInfoView.vue";
import UserInfoEditView from "@/views/user/UserInfoEditView.vue";
import CategoryView from "@/views/CategoryView.vue";
const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "home",
    component: HomeView,
  },
  {
    path: "/user/info",
    name: "用户信息",
    component: UserInfoView,
  },
  {
    path: "/auth/login",
    name: "登录",
    component: UserLoginView,
  },
  {
    path: "/auth/register",
    name: "注册",
    component: UserRegisterView,
  },
  {
    path: "/cart",
    name: "购物车",
    component: CartView,
  },
  {
    path: "/goods",
    name: "所有商品",
    component: GoodsView,
  },
  {
    path: "/good/info/:goodname",
    name: "商品详情",
    component: SingleGoodView,
  },
  {
    path: "/user/info/edit",
    name: "修改用户信息",
    component: UserInfoEditView,
  },
  {
    path: "/goods/category/:categoryname",
    name: "查看分类",
    component: CategoryView,
  },

  {
    path: "/about",
    name: "about",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/AboutView.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
