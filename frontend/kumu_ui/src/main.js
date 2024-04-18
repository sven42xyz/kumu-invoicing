import "bootstrap/dist/css/bootstrap.min.css"
import "primevue/resources/themes/aura-dark-noir/theme.css";
import "primevue/resources/themes/aura-dark-noir/fonts/Inter-roman.var.woff2";
import 'primeicons/primeicons.css'
import "bootstrap"
import "../public/css/App.css"
import "../public/aura-dark-noir/theme.css"

import { createApp } from 'vue'
import VueCookies from 'vue3-cookies'
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
import AutoComplete from 'primevue/autocomplete';
import Accordion from 'primevue/accordion';
import AccordionTab from 'primevue/accordiontab';
import AnimateOnScroll from 'primevue/animateonscroll';
import Avatar from 'primevue/avatar';
import AvatarGroup from 'primevue/avatargroup';
import Badge from 'primevue/badge';
import BadgeDirective from "primevue/badgedirective";
import BlockUI from 'primevue/blockui';
import Button from 'primevue/button';
import ButtonGroup from 'primevue/buttongroup';
import Breadcrumb from 'primevue/breadcrumb';
import Calendar from 'primevue/calendar';
import Card from 'primevue/card';
import CascadeSelect from 'primevue/cascadeselect';
import Carousel from 'primevue/carousel';
import Checkbox from 'primevue/checkbox';
import Chip from 'primevue/chip';
import Chips from 'primevue/chips';
import ColumnGroup from 'primevue/columngroup';
import ConfirmDialog from 'primevue/confirmdialog';
import ContextMenu from 'primevue/contextmenu';
import DataView from 'primevue/dataview';
import DataViewLayoutOptions from 'primevue/dataviewlayoutoptions';
import DeferredContent from 'primevue/deferredcontent';
import Dialog from 'primevue/dialog';
import DialogService from 'primevue/dialogservice'
import Divider from 'primevue/divider';
import Dock from 'primevue/dock';
import Dropdown from 'primevue/dropdown';
import DynamicDialog from 'primevue/dynamicdialog';
import FileUpload from 'primevue/fileupload';
import FloatLabel from 'primevue/floatlabel';
import FocusTrap from 'primevue/focustrap';
import Galleria from 'primevue/galleria';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import Image from 'primevue/image';
import InlineMessage from 'primevue/inlinemessage';
import Inplace from 'primevue/inplace';
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputMask from 'primevue/inputmask';
import InputNumber from 'primevue/inputnumber';
import InputOtp from 'primevue/inputotp';
import InputSwitch from 'primevue/inputswitch';
import InputText from 'primevue/inputtext';
import Knob from 'primevue/knob';
import Listbox from 'primevue/listbox';
import MegaMenu from 'primevue/megamenu';
import Menu from 'primevue/menu';
import Menubar from 'primevue/menubar';
import Message from 'primevue/message';
import MeterGroup from 'primevue/metergroup';
import MultiSelect from 'primevue/multiselect';
import OrderList from 'primevue/orderlist';
import OrganizationChart from 'primevue/organizationchart';
import Paginator from 'primevue/paginator';
import Panel from 'primevue/panel';
import PanelMenu from 'primevue/panelmenu';
import Password from 'primevue/password';
import PickList from 'primevue/picklist';
import ProgressBar from 'primevue/progressbar';
import ProgressSpinner from 'primevue/progressspinner';
import Rating from 'primevue/rating';
import RadioButton from 'primevue/radiobutton';
import Ripple from 'primevue/ripple';
import Row from 'primevue/row';
import SelectButton from 'primevue/selectbutton';
import ScrollPanel from 'primevue/scrollpanel';
import ScrollTop from 'primevue/scrolltop';
import Skeleton from 'primevue/skeleton';
import Slider from 'primevue/slider';
import Sidebar from 'primevue/sidebar';
import SpeedDial from 'primevue/speeddial';
import SplitButton from 'primevue/splitbutton';
import Splitter from 'primevue/splitter';
import SplitterPanel from 'primevue/splitterpanel';
import Stepper from 'primevue/stepper';
import StepperPanel from 'primevue/stepperpanel';
import Steps from 'primevue/steps';
import StyleClass from 'primevue/styleclass';
import TieredMenu from 'primevue/tieredmenu';
import Textarea from 'primevue/textarea';
import Toolbar from 'primevue/toolbar';
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import Tag from 'primevue/tag';
import Terminal from 'primevue/terminal';
import Timeline from 'primevue/timeline';
import ToggleButton from 'primevue/togglebutton';
import Tree from 'primevue/tree';
import TreeSelect from 'primevue/treeselect';
import TreeTable from 'primevue/treetable';
import TriStateCheckbox from 'primevue/tristatecheckbox';
import VirtualScroller from 'primevue/virtualscroller';

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
app.use(DialogService);

app.component('OverlayPanel', OverlayPanel);
app.component('ColorPicker', ColorPicker);
app.component('Field-set', Fieldset);
app.component('TabMenu', TabMenu);
app.component('DataTable', DataTable);
app.component('ColumnColumn', Column);
app.directive('tooltip', Tooltip);
app.component('ConfirmPopup', ConfirmPopup);
app.component('Toast-Toast', Toast);

app.directive('badge', BadgeDirective);
app.directive('ripple', Ripple);
app.directive('styleclass', StyleClass);
app.directive('focustrap', FocusTrap);
app.directive('animateonscroll', AnimateOnScroll);

app.component('AccorDion', Accordion);
app.component('AccordionTab', AccordionTab);
app.component('AutoComplete', AutoComplete);
app.component('AvaTar', Avatar);
app.component('AvatarGroup', AvatarGroup);
app.component('BadgeBadge', Badge);
app.component('BlockUI', BlockUI);
app.component('BreadCrumb', Breadcrumb);
app.component('ButtonButton', Button);
app.component('ButtonGroup', ButtonGroup);
app.component('CalenDar', Calendar);
app.component('CardCard', Card);
app.component('CarouSel', Carousel);
app.component('CascadeSelect', CascadeSelect);
app.component('CheckBox', Checkbox);
app.component('ChipChip', Chip);
app.component('ChipS', Chips);
app.component('ColumnColumn', Column);
app.component('ColumnGroup', ColumnGroup);
app.component('ConfirmDialog', ConfirmDialog);
app.component('ConfirmPopup', ConfirmPopup);
app.component('ContextMenu', ContextMenu);
app.component('DataTable', DataTable);
app.component('DataView', DataView);
app.component('DataViewLayoutOptions', DataViewLayoutOptions);
app.component('DeferredContent', DeferredContent);
app.component('DialogDialog', Dialog);
app.component('DividerDivider', Divider);
app.component('DockDock', Dock);
app.component('DropdownDropdown', Dropdown);
app.component('DynamicDialog', DynamicDialog);
app.component('FieldsetFieldset', Fieldset);
app.component('FileUpload', FileUpload);
app.component('FloatLabel', FloatLabel);
app.component('GalleriaGalleria', Galleria);
app.component('IconField', IconField);
app.component('ImageImage', Image);
app.component('InlineMessage', InlineMessage);
app.component('InplaceInplace', Inplace);
app.component('InputGroup', InputGroup);
app.component('InputGroupAddon', InputGroupAddon);
app.component('InputIcon', InputIcon);
app.component('InputMask', InputMask);
app.component('InputNumber', InputNumber);
app.component('InputOtp', InputOtp);
app.component('InputSwitch', InputSwitch);
app.component('InputText', InputText);
app.component('KnobKnob', Knob);
app.component('ListBox', Listbox);
app.component('MegaMenu', MegaMenu);
app.component('MenuMenu', Menu);
app.component('MenuBar', Menubar);
app.component('MessageMessage', Message);
app.component('MeterGroup', MeterGroup);
app.component('MultiSelect', MultiSelect);
app.component('OrderList', OrderList);
app.component('OrganizationChart', OrganizationChart);
app.component('OverlayPanel', OverlayPanel);
app.component('PagiNator', Paginator);
app.component('PanelPanel', Panel);
app.component('PanelMenu', PanelMenu);
app.component('PassWord', Password);
app.component('PickList', PickList);
app.component('ProgressBar', ProgressBar);
app.component('ProgressSpinner', ProgressSpinner);
app.component('RadioButton', RadioButton);
app.component('RatingRating', Rating);
app.component('RowRow', Row);
app.component('SelectButton', SelectButton);
app.component('ScrollPanel', ScrollPanel);
app.component('ScrollTop', ScrollTop);
app.component('SliderSlider', Slider);
app.component('SideBar', Sidebar);
app.component('SkeletonSkeleton', Skeleton);
app.component('SpeedDial', SpeedDial);
app.component('SplitButton', SplitButton);
app.component('SplitTer', Splitter);
app.component('SplitterPanel', SplitterPanel);
app.component('StepPer', Stepper);
app.component('StepperPanel', StepperPanel);
app.component('StepsSteps', Steps);
app.component('TabMenu', TabMenu);
app.component('TabView', TabView);
app.component('TabPanel', TabPanel);
app.component('TagTag', Tag);
app.component('TextArea', Textarea);
app.component('TermiNal', Terminal);
app.component('TieredMenu', TieredMenu);
app.component('TimeLine', Timeline);
app.component('ToolBar', Toolbar);
app.component('ToggleButton', ToggleButton);
app.component('TreeTree', Tree);
app.component('TreeSelect', TreeSelect);
app.component('TreeTable', TreeTable);
app.component('TriStateCheckbox', TriStateCheckbox);
app.component('VirtualScroller', VirtualScroller);

app.config.compilerOptions.isCustomElement = (tag) => {
    return tag.startsWith('h7-')
}
app.mount('#app');
