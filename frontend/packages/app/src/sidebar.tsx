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
    <SidebarItem icon={CreateComponentIcon} to="diagnosis" text="Diagnosis" />
    <SidebarItem icon={CreateComponentIcon} to="triageresult" text="TriageResult" />
    <SidebarItem icon={CreateComponentIcon} to="medicalprocedure" text="MedicalProcedre" />
    <SidebarItem icon={CreateComponentIcon} to="diagnosis" text="Diagnosis" />
    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarThemeToggle />
    <SidebarUserSettings />
    <SidebarPinButton />
  </Sidebar>
);