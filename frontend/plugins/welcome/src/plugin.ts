import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import DiagnosisPage from './components/DiagnosisPage';
import LoginPage from './components/login';
import TriageResult from './components/TriageResult';
import Result from './components/Result';
import MedicalProcedure from './components/Medicalprocedure';
import Medicalresults from './components/Medicalresults';
import Table from './components/Table';
import righttotreatment from './components/Create';
import Righttable from './components/Righttable';
import PatientRegistration from './components/PatientRegistration';
import PatientRegistrationTables from './components/PatientRegistrationTable';
import CreateAppointment from './components/CreateAppointment';
import AppointmentResults from './components/AppointmentResults';
import DiagnosisTables from './components/DiagnosisTable';
import SearchDiagnosisPage from './components/SearchDiagnosisPage';
import { Cookies } from './Cookie'

var ck = new Cookies()
var role = ck.GetRole()

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
    router.registerRoute('/Create', righttotreatment);
    router.registerRoute('/Righttable', Righttable);
    router.registerRoute('/patientregistration', PatientRegistration);
    router.registerRoute('/patientregistrationtable', PatientRegistrationTables);
    router.registerRoute('/createappointment', CreateAppointment);
    router.registerRoute('/appointmentresults', AppointmentResults);
    router.registerRoute('/diagnosistable', DiagnosisTables);
    router.registerRoute('/searchdiagnosisPage', SearchDiagnosisPage);
  },
});