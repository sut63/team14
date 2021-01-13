import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Group14 from './components/Group14';
import Personalpage from './components/Personalpage';
import Tablepersonal from './components/Tablepersonal';

 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/Group14', Group14);
    router.registerRoute('/Personalpage', Personalpage);
    router.registerRoute('/Tablepersonal', Tablepersonal);
  },
});