import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import DiagnosisPage from './components/DiagnosisPage';
import LoginPage from './components/login';
import TriageResult from './components/TriageResult';
import Result from './components/Result';
import MedicalProcedure from './components/Medicalprocedure';
import Medicalresults from './components/Medicalresults';
import Table from './components/Table';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/welcome', WelcomePage);
    router.registerRoute('/diagnosis', DiagnosisPage);
    router.registerRoute('/', LoginPage);
    router.registerRoute('/triageresult', TriageResult);
    router.registerRoute('/result', Result);
    router.registerRoute('/medicalprocedure', MedicalProcedure);
    router.registerRoute('/medicalresults', Medicalresults);
    router.registerRoute('/table', Table);
  },
});