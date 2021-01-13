import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import DiagnosisPage from './components/DiagnosisPage';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/welcome', WelcomePage);
    router.registerRoute('/diagnosis', DiagnosisPage);
  },
});

