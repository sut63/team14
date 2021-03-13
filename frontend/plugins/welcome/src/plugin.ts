import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Group14 from './components/Group14';
import CreateNewCustomer from './components/Customers';
import CustomerTable from './components/CustomerTables';
import Productcreates from './components/Productcreates';
import Producttable from './components/Producttable';
import Fixpage from './components/Fixpage';
import FixSearch from './components/FixSearch';
import Tablefix from './components/Tablefix';
import AdminrepairUI from './components/AdminrepairUI';
import AdminrepairSearch from './components/AdminrepairSearch';
import Selectadminrepair from './components/Selectadminrepair';
import Receiptcreate from './components/Createreceipts';
import Tablereceipt from './components/Tablereceipt';
import CustomerSearch from './components/CustomerSearch';
import Personalcreate from './components/Personalcreate/';
import Personaltable from './components/Personaltable';
import Personalsearch from './components/Personalsearch';
import ProductSearch from './components/ProductSearch';
import ReceiptSearch from './components/ReceiptSearch';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/Group14', Group14);
    router.registerRoute('/Personalcreate', Personalcreate);
    router.registerRoute('/Personaltable', Personaltable);
    router.registerRoute('/Personalsearch', Personalsearch);
    router.registerRoute('/createnewcustomer', CreateNewCustomer);
    router.registerRoute('/customertable', CustomerTable);
    router.registerRoute('/Producttables', Producttable);
    router.registerRoute('/Productcreate', Productcreates);
    router.registerRoute('/ProductSearch', ProductSearch);
    router.registerRoute('/Fixpage', Fixpage);
    router.registerRoute('/FixSearch', FixSearch);
    router.registerRoute('/Tablefix', Tablefix);
    router.registerRoute('/AdminrepairUI', AdminrepairUI);
    router.registerRoute('/AdminrepairSearch', AdminrepairSearch);
    router.registerRoute('/Selectadminrepair', Selectadminrepair);
    router.registerRoute('/createreceipt', Receiptcreate);
    router.registerRoute('/CustomerSearch', CustomerSearch);
    router.registerRoute('/Tablereceipt', Tablereceipt);
    router.registerRoute('/ReceiptSearch', ReceiptSearch);

  },
});