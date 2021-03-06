import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="./" text="Home" />
    <SidebarItem icon={CreateComponentIcon} to="patientregistration" text="Patient Registration" />
    <SidebarItem icon={CreateComponentIcon} to="Create" text="Right to treatment" />
    <SidebarItem icon={CreateComponentIcon} to="triageresult" text="TriageResult" />
    <SidebarItem icon={CreateComponentIcon} to="loginpage" text="Diagnosis" />
    <SidebarItem icon={CreateComponentIcon} to="medicalprocedure" text="MedicalProcedure" />
    <SidebarItem icon={CreateComponentIcon} to="createappointment" text="AppointmentResults" />
   
    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarThemeToggle />
    <SidebarUserSettings />
    <SidebarPinButton />
  </Sidebar>
);