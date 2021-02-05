import React, { FC } from 'react';
import {
  Typography,
  Grid,
  List,
  ListItem,
  ListItemText,
  Link,
} from '@material-ui/core';
import {
  Content,
  InfoCard,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';

const WelcomePage: FC<{}> = () => {

  return (
    <Page theme={pageTheme.home}>
      <Header
        title="ยินตีต้อนรับสู่ระบบประวัติผู้ป่วย"
      >
      </Header>
      <Content>
        <ContentHeader title="ระบบประวัติผู้ป่วย">
        </ContentHeader>
        <Grid container>
          <Grid item xs={12} md={6}>
            <InfoCard>
              <Typography variant="body1" gutterBottom>
                &nbsp; &nbsp; &nbsp; ระบบประวัติผู้ป่วย เป็นระบบที่ผู้ใช้ระบบบสามารถตรวจสอบและแก้ไขข้อมูลผู้ป่วย
                และสามารถเพิ่มข้อมูลของผู้ป่วยที่ยังไม่มีในระบบได้
                สามารถตรวจสอบสิทธิการรักษา ผลการคัดกรอง ผลการตรวจรักษา ดูบันทึกการทำหัตการ
                และยังมีระบบนัดหมายผู้ป่วยอีกด้วย
              </Typography>
              <br></br>
              <Typography variant="h6" gutterBottom>
                สมาชิกในทีม
              </Typography>
              <List>
                <ListItem>
                  <ListItemText primary="B6102852 นายฤชากร กลิ่นสุข - ระบบนัดหมาย ระบบค้นหาการนัดหมาย" />
                </ListItem>
                <ListItem>
                  <ListItemText primary="B6104320 นายวุฒิศักดิ์ คุชิตา - ระบบลงทะเบียนผู้ป่วย ระบบค้นหาข้อมูลผู้ป่วย" />
                </ListItem>
                <ListItem>
                  <ListItemText primary="B6104696 นายสมเกียรติ จบสูงเนิน - ระบบบันทึกการคัดกรอง ระบบค้นหาบันทึกการคัดกรอง" />
                </ListItem>
                <ListItem>
                  <ListItemText primary="B6109189 นางสาวสุภาพร บุญอิทร์ - ระบบสิทธิการรักษา ระบบค้นหาสิทธิการรักษา" />
                </ListItem>
                <ListItem>
                  <ListItemText primary="B6116262 นางสาวปานตา เสาวภา - ระบบบันทึกการทำหัตถการ ระบบค้นหาการทำหัตถการ" />
                </ListItem>
                <ListItem>
                  <ListItemText primary="B6117368 นายนพชัย อัตถาวงศ์ - ระบบการตรวจรักษา ระบบค้นหาประวัติการตรวจรักษา" />
                </ListItem>
              </List>
            </InfoCard>
          </Grid>
          <Grid item>
            <InfoCard>
              <Typography variant="h5">ระบบย่อย #1</Typography>
              <List>
                <ListItem>
                  <Link href="http://localhost:3000/patientregistration">ระบบลงทะเบียนผู้ป่วย</Link>
                </ListItem>
                <ListItem>
                  <Link href="http://localhost:3000/Loginright">ระบบบันทึกสิทธิการรักษา</Link>
                </ListItem>
                <ListItem>
                  <Link href="http://localhost:3000/triageresult">ระบบบันทึกการคัดกรอง</Link>
                </ListItem>
                <ListItem>
                  <Link href="http://localhost:3000/diagnosis">ระบบการตรวจรักษา</Link>
                </ListItem>
                <ListItem>
                  <Link href="http://localhost:3000/medicalprocedure">ระบบบันทึกการทำหัตถการ</Link>
                </ListItem>
                <ListItem>
                  <Link href="http://localhost:3000/loginappointment">ระบบนัดหมาย</Link>
                </ListItem>
              </List>
            </InfoCard>
          </Grid>
          <Grid item>
            <InfoCard>
              <Typography variant="h5">ระบบย่อย #2</Typography>
              <List>
                <ListItem>
                  <Link href="http://localhost:3000/searchpatient">ระบบค้นหาข้อมูลผู้ป่วย</Link>
                </ListItem>
                <ListItem>
                  <Link href="http://localhost:3000/Searchrighttotreatment">ระบบค้นหาสิทธิการรักษา</Link>
                </ListItem>
                <ListItem>
                  <Link href="http://localhost:3000/triageresult">ระบบค้นหาบันทึกการคัดกรอง (ลิงก์ยังไม่ถูก)</Link>
                </ListItem>
                <ListItem>
                  <Link href="http://localhost:3000/SearchdiagnosisPage">ระบบค้นหาประวัติการตรวจรักษา</Link>
                </ListItem>
                <ListItem>
                  <Link href="http://localhost:3000/medicalprocedure">ระบบค้นหาบันทึกการคัดกรอง (ลิงก์ยังไม่ถูก)</Link>
                </ListItem>
                <ListItem>
                  <Link href="http://localhost:3000/createappointment">ระบบค้นหาการนัดหมาย (ลิงก์ยังไม่ถูก)</Link>
                </ListItem>
              </List>
            </InfoCard>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
