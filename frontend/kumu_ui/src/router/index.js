import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '../components/vues/LandingPage.vue'
import DashBoard from '../components/vues/DashBoard.vue'

const routes = [
    { path: '/', component: LandingPage },
    { path: '/dashboard', name: 'dashboard', component: DashBoard },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;