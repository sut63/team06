import React, { useState, useEffect } from 'react';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, TextField, Button, Grid } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
import Swal from 'sweetalert2'
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';

import { EntHospital } from '../../api/models/EntHospital';
import { EntPatient } from '../../api/models/EntPatient';
import { EntRightToTreatment } from '../../api/models/EntRightToTreatment';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(2),
    },
    insideLabel: {
      margin: theme.spacing(1),
    },
    button: {
      marginLeft: '50px',
    },
    textField: {
      width: 600,
      marginLeft: 10,
      marginRight: -10,
    },
    select: {
      width: 600,
      marginLeft: 10,
      marginRight: -10,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    pa: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
  }),
);

const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});


export default function ComponentsTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);

  const [checkvalue, setCheckValue] = useState(false);

  const [checkPatientName, setPatientNames] = useState(false);
  const [hospitals, setHospitals] = useState<EntHospital[]>([]);
  const [patient, setPatient] = useState<EntPatient[]>([]);
  const [righttotreatment, setRightToTreatments] = useState<EntRightToTreatment[]>([]);

  const [patientname, setPatientName] = useState(String);
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getPatients = async () => {
      const res = await http.listPatient({ offset: 0 });
      setLoading(false);
      setPatient(res);
    };
    getPatients();
    const getRighttable = async () => {
      const res = await http.listRighttotreatment({ offset: 0 });
      setLoading(false);
      setRightToTreatments(res);
    };
    getRighttable();
  }, [loading]);

  const Patienthandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setPatientNames(false);
    setPatientName(event.target.value as string);

  };

  const cleardata = () => {
    setPatientName("");
    setSearch(false);
    setPatientNames(false);
    setSearch(false);

  }

  const checkresearch = async () => {
    var check = false;
    righttotreatment.map(item => {
      if (patientname != "") {
        if (item.edges?.patient?.patientName?.startsWith(patientname)) {
          setPatientNames(true);
          alertMessage("success", "ค้นหาข้อมูลสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูลที่ค้นหา");
    }
    console.log(checkPatientName)
    if (patientname == "") {
      alertMessage("info", "กรอกข้อมูลชื่อผู้ป่วย");
    }
  };

  return (

    <Page theme={pageTheme.home}>
      <Header title="ระบบค้นหาสิทธิการรักษาของผู้ป่วย">
      </Header>
      <Content>
        <ContentHeader title="ดูข้อมูลสิทธิการรักษา">
          <Button
            href="/Create"
            variant="contained"
            color="primary"
          >
            ย้อนกลับ
          </Button>
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
              size="small"
            >
              <div className={classes.paper}><strong>กรอกชื่อผู้ป่วยเพื่อค้นหา</strong></div>
              <TextField
                style={{ width: 500, marginLeft: 7, marginRight: -7, marginTop: 5 }}
                className={classes.textField}
                id="searchrighttotreatmentt"
                variant="outlined"
                color="primary"
                type="string"
                size="small"
                value={patientname}
                onChange={Patienthandlehange}
              />
            </FormControl>
          </form>
        </div>
        <div className={classes.root}>
          <Button
            onClick={() => {
              checkresearch();
              setSearch(true);
            }}
            variant="contained"
            color="secondary"
            startIcon={<SearchTwoToneIcon />}
          >
            ค้นหาข้อมูล
          </Button>
          <div>&nbsp;&nbsp;&nbsp;</div>

        </div>

        <br></br>

        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>
              {search ? (
                <div>
                  {  checkPatientName ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                          <TableCell align="center">No.</TableCell>
           <TableCell align="center">ผู้ป่วย</TableCell>
           <TableCell align="center">เลขบัตรประจำตัวประชาชน</TableCell>
           <TableCell align="center">อายุ</TableCell>
           <TableCell align="center">ประเภทสิทธิการรักษา</TableCell>
           <TableCell align="center">โรงพยาบาลที่ใช้สิทธิ</TableCell>
           <TableCell align="center">เบอร์โทรศัพท์</TableCell>
           <TableCell align="center">วันที่เริ่มใช้สิทธิ</TableCell>
           <TableCell align="center">วันที่หมดสิทธิ</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {righttotreatment.filter((filter: any) => filter.edges?.patient?.patientName.startsWith(patientname)).map((item: any) => (
                            <TableRow key={item.id}>
                            <TableCell align="center">{item.id}</TableCell>
                            <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                            <TableCell align="center">{item.idennum}</TableCell>
                            <TableCell align="center">{item.age}</TableCell>
                            <TableCell align="center">{item.edges?.rightToTreatmentType?.typeName}</TableCell>
                            <TableCell align="center">{item.edges?.hospital.hospitalName}</TableCell>
                            <TableCell align="center">{item.tel}</TableCell>
                            <TableCell align="center">{moment(item.startTime).format('DD/MM/YYYY')}</TableCell>
                            <TableCell align="center">{moment(item.endTime).format('DD/MM/YYYY')}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : patientname == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                              <TableCell align="center">No.</TableCell>
           <TableCell align="center">ผู้ป่วย</TableCell>
           <TableCell align="center">เลขบัตรประจำตัวประชาชน</TableCell>
           <TableCell align="center">อายุ</TableCell>
           <TableCell align="center">ประเภทสิทธิการรักษา</TableCell>
           <TableCell align="center">โรงพยาบาลที่ใช้สิทธิ</TableCell>
           <TableCell align="center">เบอร์โทรศัพท์</TableCell>
           <TableCell align="center">วันที่เริ่มใช้สิทธิ</TableCell>
           <TableCell align="center">วันที่หมดสิทธิ</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>

                              {righttotreatment.map((item: any) => (
                                <TableRow key={item.id}>
                                  <TableCell align="center">{item.id}</TableCell>
                                  <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                                  <TableCell align="center">{item.idennum}</TableCell>
                                  <TableCell align="center">{item.age}</TableCell>
                                  <TableCell align="center">{item.edges?.rightToTreatmentType?.typeName}</TableCell>
                                  <TableCell align="center">{item.edges?.hospital.hospitalName}</TableCell>
                                  <TableCell align="center">{item.tel}</TableCell>
                                  <TableCell align="center">{moment(item.startTime).format('DD/MM/YYYY')}</TableCell>
                                  <TableCell align="center">{moment(item.endTime).format('DD/MM/YYYY')}</TableCell>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableContainer>

                      </div>
                    ) : null}
                </div>
              ) : null}
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}