// 定义请求参数和响应数据的接口
interface LoginParams {
  username: string;
  password: string;
}

interface LoginResponse {
  token: string;
  userId: string;
}

interface RegisterParams {
  username: string;
  password: string;
  email: string;
}

interface RegisterResponse {
  message: string;
}

interface GoodsParams {
  keyword?: string;
  category?: string;
  page?: number;
  limit?: number;
}

interface GoodsResponse {
  data: Good[];
  total: number;
}

interface Good {
  category: string;
  name: string;
  price: number;
  description: string;
}

interface CommentParams {
  good_id: string;
  content: string;
}

interface CommentResponse {
  data: Comment[];
}

interface CategoryParams {
  name: string;
}

interface CategoryResponse {
  data: Category[];
}

interface Category {
  id: string;
  name: string;
}

interface CartParams {
  good_id: string;
  quantity: number;
}

interface CartResponse {
  data: CartItem[];
}

interface CartItem {
  goodname: string;
  username: string;
  quantity: number;
  price: number;
}

interface User {
  username: string;
  email: string;
}
