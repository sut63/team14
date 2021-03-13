import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';

import { EntCustomer } from '../../api/models/EntCustomer';

import Swal from 'sweetalert2'
import { Link as RouterLink } from 'react-router-dom';
import moment from 'moment';
import { Page, pageTheme, Header, Content, Link, ContentHeader } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';

import { styled } from '@material-ui/core/styles';
import { compose, spacing, palette, sizing, shadows   } from '@material-ui/system';
import { Alert } from '@material-ui/lab';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    headsearch: {
      width: 'auto',
      margin: '10px',
      color: '#FFFFFF',
      background: '#2196F3',
    },
    margin: {
      margin: theme.spacing(1),
    },
    margins: {
      margin: theme.spacing(2),
    },

    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
    paper: {
      marginTop: theme.spacing(3),
      marginBottom: theme.spacing(3),
    },
    table: {
      minWidth: 500,
    },


  }),
);
const Toast = Swal.mixin({
  // toast: true,
  position: 'center',
  showConfirmButton: false,
  //timer: 3000,
  //timerProgressBar: true,
  showCloseButton: true,

});


export default function ComponentsTable() {

    //--------------------------
    ;

  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);

  //---------------------------
  const [checkcustomername, setCustomernames] = useState(false);
  const [customer, setCustomer] = useState<EntCustomer[]>([])

  //--------------------------
  const [cmname, setcmname] = useState(String);
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบค้นหาข้อมูลลูกค้า' };

  const Box = styled('div')(compose(spacing, palette, shadows, sizing ));
  
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

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

 
  //-------------------
  const customernamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setcmname(event.target.value as string);
  };

  const cleardata = () => {
    setcmname("");
    setStatus(false)
    setCustomer([]);
  }


  const SearchCustomer = async () => {
    setStatus(true);
    setAlert(true);
    const apiUrl = `http://localhost:8080/api/v1/searchcustomers?customer=${cmname}`;
    const requestOptions = {
      method: 'GET',
    };
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.data)
        setStatus(true);
        setAlert(false);
        setCustomer([]);
        if (data.data != null) {
          if (data.data.length >= 1) {
            setStatus(true);
            setAlert(true);
            console.log(data.data)
            setCustomer(data.data);
          }
        }
      });
    }

  return (

    <Page theme={pageTheme.tool}>
      <Header title= "Customer System" type="Computer Repair System" >

      </Header> 
      <Content >
        <ContentHeader title="ค้นหาข้อมูลลูกค้า">
        {status ? (
              <div>
                {alert ? (
                  <Alert severity="success">
                    ค้นหาข้อมูลลูกค้าสำเร็จ
                  </Alert>
                )
                  : (
                    <Alert severity="warning" style={{ marginTop: 20 }}>
                      ค้นหาข้อมูลลูกค้าไม่สำเร็จ
                    </Alert>
                  )}
              </div>
            ) : null}
        </ContentHeader>

        <Grid container item xs={12} justify="center">
          <Grid item xs={5}>
            <Paper>
            
              <Typography align="center" >
                <div style={{ background: 'linear-gradient(45deg, #CCCCCC 15%, #CCCCCC 120%)', height: 45 }}>
                  <h1 style={
                    {
                      //color: "#000000",
                      //borderRadius: 5,
                      //height: 18,
                      //padding: '0 30px',
                      fontSize: '25px',
                    }}>
                    ค้นหาข้อมูลลูกค้า
            </h1>
                </div>

                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}><strong>กรุณากรอกชื่อลูกค้าที่ต้องการค้นหา</strong></div>
                    <TextField
                      id="customername"
                      value={cmname}
                      onChange={customernamehandlehange}
                      type="string"
                      size="small"

                      style={{ width: 250 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    SearchCustomer();
                  }}
                  
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#5C9DC0", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 10px',

                      }
                    }>
                    ค้นหาข้อมูล
            </h3>
                </Button>
                <Button
                  onClick={() => {
                    cleardata();

                  }}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#DD0000", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 25px',

                      }
                    }>
                    ลบ
            </h3>
                </Button>
              </Typography>
            </Paper>
          </Grid>
        </Grid>
        
        
        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>
              <TableContainer component={Paper}>
                <Table className={classes.table} aria-label="simple table">
                  <TableHead>
                    <TableRow>
                      <TableCell align="center">ลำดับที่</TableCell>
                      <TableCell align="center">คำนำหน้าชื่อ</TableCell>
                      <TableCell align="center">ชื่อ-นามสกุล</TableCell>
                      <TableCell align="center">ที่อยู่</TableCell>
                      <TableCell align="center">เบอร์โทร</TableCell>
                      <TableCell align="center">เลขบัตรประจำตัวประชาชน</TableCell>
                      <TableCell align="center">เพศ</TableCell>
                      
                    </TableRow>
                  </TableHead>
                <TableBody>
                {customer.map((item: any) => (
                  <TableRow key={item.id}>
                    <TableCell align="center">{item.id}</TableCell>
                    <TableCell align="center">{item.edges?.Title?.titlename}</TableCell>
                    <TableCell align="center">{item.Customername}</TableCell>
                    <TableCell align="center">{item.Address}</TableCell>
                    <TableCell align="center">{item.Phonenumber}</TableCell>
                    <TableCell align="center">{item.Identificationnumber}</TableCell>
                    <TableCell align="center">{item.edges?.Gender?.Gendername}</TableCell>
                    
                  </TableRow>
                ))}
                </TableBody>
              </Table>
            </TableContainer>
            </Paper>
        </Grid>
        </Grid>

        <div>&nbsp;&nbsp;&nbsp;</div>
        <Typography align="center" >
        <div>&nbsp;&nbsp;&nbsp;</div>
        <Button 
        style={{ marginLeft: 20 }} 
        href="/CustomerTable"
        variant="contained"  
        > 
        กลับหน้าตารางข้อมูลลูกค้า 
        </Button>

        <Button 
        style={{ marginLeft: 20 }} 
        href="/Group14"
        variant="contained"  
        > 
        กลับหน้าหลัก 
        </Button>
        </Typography>
        <div>&nbsp;&nbsp;&nbsp;</div>

      </Content>
    </Page>
  );

}