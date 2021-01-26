import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { Button, TableCell,} from '@material-ui/core';
import { Content,Header,Page,pageTheme,ContentHeader, AlertDisplay } from '@backstage/core';
import Tables from './Tables';

const TriageResultSearch: FC<{}> = () => {

  let name = "aaaaa";

 return (
   <Page theme={pageTheme.home}>
     <Header
       title={'ผลการคัดกรองผู้ป่วย'}
     >
       <TableCell align={
        "center"} >
          <text> {name} </text>
        <br/><br/>
        <Button
            
            component={RouterLink}
            to="/triageresultlogin"
            variant="contained"
            color="secondary"
        >
            Logout
        </Button>
    </TableCell>
     </Header>

     <Content>
       <ContentHeader title="">
             <Button 
             component={RouterLink} 
             to='/triageresult'
             variant="contained"
             color="primary" >
                 ADD Screening
             </Button>
       </ContentHeader>
            <Tables></Tables>
     </Content>
   </Page>
 );
};

export default TriageResultSearch;