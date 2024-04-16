import "bootstrap/dist/css/bootstrap.min.css"
import "primevue/resources/themes/aura-light-noir/theme.css";
import "primevue/resources/themes/aura-dark-noir/theme.css";
import "primevue/resources/themes/aura-light-noir/fonts/Inter-roman.var.woff2";
import "primevue/resources/themes/aura-dark-noir/fonts/Inter-roman.var.woff2";
import 'primeicons/primeicons.css'
import "bootstrap"
import "../public/css/App.css"
import "../public/aura-dark-noir/theme.css"

import { createApp } from 'vue'
import VueCookies from 'vue3-cookies'
//import VueSocketIO from 'vue-socket.io'
import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config';
import Fieldset from 'primevue/fieldset';
import ColorPicker from "primevue/colorpicker"
import Tooltip from 'primevue/tooltip';
import ConfirmPopup from 'primevue/confirmpopup';
import TabMenu from 'primevue/tabmenu';
import ConfirmationService from 'primevue/confirmationservice';
import OverlayPanel from 'primevue/overlaypanel';
import Toast from 'primevue/toast';
import ToastService from 'primevue/toastservice';

import DataTable from 'primevue/datatable';
import Column from 'primevue/column';


const app = createApp(App);


app.use(router);
app.use(VueCookies, {
    expireTimes: '0',
    sameSite: 'Strict'
});
app.use(
 //   new VueSocketIO({
  //      connection: process.env.VUE_APP_SOCKET_ENDPOINT
 //   })
);
app.use(PrimeVue, { ripple: true, inputStyle: "outlined" });
app.use(ToastService);
app.use(ConfirmationService);
app.component('OverlayPanel', OverlayPanel);
app.component('ColorPicker', ColorPicker);
app.component('Field-set', Fieldset);
app.component('TabMenu', TabMenu);
app.component('DataTable', DataTable);
app.component('ColumnColumn', Column);
app.directive('tooltip', Tooltip);
app.component('ConfirmPopup', ConfirmPopup);
app.component('Toast-Toast', Toast);
app.config.compilerOptions.isCustomElement = (tag) => {
    return tag.startsWith('h7-')
}
app.mount('#app');
