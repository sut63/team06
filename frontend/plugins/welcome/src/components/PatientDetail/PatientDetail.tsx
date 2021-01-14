import React, { FC, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import SaveIcon from '@material-ui/icons/Save';
import AccountBoxIcon from '@material-ui/icons/AccountBox';
import Swal from 'sweetalert2';

import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
} from '@backstage/core';

import {
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  Button,
  Box,
} from '@material-ui/core';

import {
  DefaultApi,
  EntGender,
  EntPrefix,
  EntBloodType,
  EntPatient,
} from '../../api';

const useStyles = makeStyles({
  boxStyle: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    flexDirection: 'row'
  },
  formControl: {
    width: 300,
    height: 65,
  },
});

interface patientDetail {
  prefix: number;
  gender: number;
  bloodtype: number;
  patient: number;
}

const PatientDetail: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [patient_detail, setPatientDetail] = React.useState<
    Partial<patientDetail>
  >({});

  const [genders, setGenders] = React.useState<EntGender[]>([]);
  const [bloodtypes, setBloodtypes] = React.useState<EntBloodType[]>([]);
  const [prefixs, setPrefixs] = React.useState<EntPrefix[]>([]);
  const [patients, setPatients] = React.useState<EntPatient[]>([]);


  // alert setting
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


  const getGenders = async () => {
    const res = await http.listGender({ limit: 10, offset: 0 });
    setGenders(res);
  };

  const getBloodtypes = async () => {
    const res = await http.listBloodtype({ limit: 10, offset: 0 });
    setBloodtypes(res);
  };

  const getPrefixs = async () => {
    const res = await http.listPrefix({ limit: 10, offset: 0 });
    setPrefixs(res);
  };

  const getPatients = async () => {
    const res = await http.listPatient({ limit: 10, offset: 0 });
    setPatients(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getGenders();
    getBloodtypes();
    getPrefixs();
    getPatients();
  }, []);

  // set data to object patient_detail
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof PatientDetail;
    const { value } = event.target;
    setPatientDetail({ ...patient_detail, [name]: value });
  };

  // clear input form
  function clear() {
    setPatientDetail({});
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/patient-details';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(patient_detail),
    };


    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.patient === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  return (

    <Page theme={pageTheme.home}>
      <Header
        title="ระบบลงทะเบียนผู้ป่วย"
        subtitle="สามารถลงทะเบียนและเพิ่มข้อมูลผู้ป่วยได้ที่นี่"
      ></Header>

      <Content>
        <ContentHeader title="เพิ่มรายละเอียดผู้ป่วย">
          <Box
            display="flex"
            justifyContent="flex-start"
            flexDirection="column"
          >
            <AccountBoxIcon style={{ fontSize: 100 }}></AccountBoxIcon>

            <Link component={RouterLink} to="/" >
              ออกจากระบบ
                        </Link>
          </Box>

        </ContentHeader>

        <Grid container>
          <Grid item xs={12}>

          <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ชื่อผู้ป่วย</InputLabel>
                <Select
                  name="patient"
                  value={patient_detail.patient || ''}
                  onChange={handleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.patientName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Box>

            <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>คำนำหน้า</InputLabel>
                <Select
                  name="prefix"
                  value={patient_detail.prefix || ''}
                  onChange={handleChange}
                >
                  {prefixs.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.prefixValue}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Box>
            <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เพศ</InputLabel>
                <Select
                  name="gender"
                  value={patient_detail.gender || ''}
                  onChange={handleChange}
                >
                  {genders.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.genderValue}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Box>
            <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>หมู่เลือด</InputLabel>
                <Select
                  name="bloodtype"
                  value={patient_detail.bloodtype || ''}
                  onChange={handleChange}
                >
                  {bloodtypes.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.bloodValue}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Box>

            <Box display="flex" justifyContent="center" >
              <Button
                style={{ width: 140, height: 40 }}
                variant="contained"
                color="primary"
                startIcon={<SaveIcon />}
                onClick={save}>
                บันทึกข้อมูล
                            </Button>
              <Button
                style={{ width: 140, height: 40, marginLeft: 20 }}
                component={RouterLink}
                to="/PatientRegistrationTable"
                variant="contained"
              >
                ดูข้อมูล
                            </Button>

            </Box>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};

export default PatientDetail;