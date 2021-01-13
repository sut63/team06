import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import diagnosisPage from './components/diagnosisPage';
import { Cookies } from './Cookies';


var ck = new Cookies()
var role = ck.GetRole()

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/welcome', WelcomePage);
    router.registerRoute('/diagnosis', diagnosisPage);
  },
});




