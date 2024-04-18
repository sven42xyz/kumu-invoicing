import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '../components/vues/LandingPage.vue'
import DashBoard from '../components/vues/DashBoard.vue'
import InvoiceUpload from '../components/vues/InvoiceUpload.vue'

const routes = [
    { path: '/', component: LandingPage },
    { path: '/dashboard', name: 'dashboard', component: DashBoard },
    { path: '/inbound', name: 'inbound', component: InvoiceUpload },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;