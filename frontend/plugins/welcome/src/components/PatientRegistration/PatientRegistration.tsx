import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import {
  Grid,
  Box,
} from '@material-ui/core';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
      display: 'flex',
      flexWrap: 'wrap',
      width: '50ch',
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);

const initialPatientState = {};

export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [patient, setPatient] = useState(initialPatientState);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const handleInputChange = (event: any) => {
    const { id, value } = event.target;
    setPatient({ ...patient, [id]: value });
  };

  const CreatePatient = async () => {
    const res = await api.createPatient({ patient });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title="ระบบลงทะเบียนผู้ป่วย"
        subtitle="สามารถลงทะเบียนและเพิ่มข้อมูลผู้ป่วยได้ที่นี่"
      >
        <Button
          component={RouterLink}
          to="/"
          variant="contained"
          color="secondary"
        >
          Logout
        </Button>

      </Header>
      <Content>
        <ContentHeader title="ลงทะเบียนผู้ป่วย">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  ลงทะเบียนสำเร็จ!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    ลงทะเบียนไม่สำเร็จ!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            >
              <TextField
                id="hospitalNumber"
                label="หมายเลขผู้ป่วย"
                variant="outlined"
                type="string"
                size="medium"
                value={patient.hospitalNumber}
                onChange={handleInputChange}
              />
            </FormControl>

            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            >
              <TextField
                id="patientName"
                label="ชื่อผู้ป่วย"
                variant="outlined"
                type="string"
                size="medium"
                value={patient.patientName}
                onChange={handleInputChange}
              />
            </FormControl>

            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            >
              <TextField
                id="drugAllergy"
                label="ประวัติการแพ้ยา"
                variant="outlined"
                type="string"
                size="medium"
                value={patient.drugAllergy}
                onChange={handleInputChange}
              />
            </FormControl>

            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreatePatient();
                }}
                style={{ marginLeft: 35 }}
                variant="contained"
                color="primary"
              >
                ลงทะเบียนผู้ป่วย
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/PatientDetail"
                variant="contained"
              >
                เพิ่มรายละเอียดของผู้ป่วย
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}