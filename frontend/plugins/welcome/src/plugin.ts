import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Group14 from './components/Group14';
import Personalpage from './components/Personalpage';
import Tablepersonal from './components/Tablepersonal';
import CreateNewCustomer from './components/Customers';
import CustomerTable from './components/CustomerTables';
import Productcreates from './components/Productcreates';
import Producttable from './components/Producttable';
import Fixpage from './components/Fixpage';
import Tablefix from './components/Tablefix';

 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/Group14', Group14);
    router.registerRoute('/Personalpage', Personalpage);
    router.registerRoute('/Tablepersonal', Tablepersonal);
    router.registerRoute('/createnewcustomer', CreateNewCustomer);
    router.registerRoute('/customertable', CustomerTable);
    router.registerRoute('/Producttables', Producttable);
    router.registerRoute('/Productcreate', Productcreates);
    router.registerRoute('/Fixpage', Fixpage);
    router.registerRoute('/Tablefix', Tablefix);
  },
});