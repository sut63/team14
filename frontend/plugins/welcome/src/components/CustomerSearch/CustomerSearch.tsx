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
import { Page, pageTheme, Header, Content, Link } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';


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

  //---------------------------
  const [checkcustomername, setCustomernames] = useState(false);
  const [customer, setCustomer] = useState<EntCustomer[]>([])

  //--------------------------
  const [customername, setCustomername] = useState(String);
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบค้นหาข้อมูลลูกค้า' };
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getCustomers = async () => {
      const res = await api.listCustomer({ offset: 0 });
      setLoading(false);
      setCustomer(res);
    };
    getCustomers();
  }, [loading]);

  //-------------------
  const customernamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setCustomernames(false);
    setCustomername(event.target.value as string);

  };

  const cleardata = () => {
    setCustomername("");
    setSearch(false);
    setCustomernames(false);
    setSearch(false);

  }
  //---------------------
  const checkresearch = async () => {
    var check = false;
    customer.map(item => {
      if (customername != "") {
        if (item.customername?.includes(customername)) {
          setCustomernames(true);
          alertMessage("success", "พบข้อมูลที่ค้นหา");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูลที่ค้นหา");
    }
    console.log(checkcustomername)
    if (customername == "") {
      alertMessage("info", "แสดงข้อมูลลูกค้าทั้งหมดทั้งหมดในระบบ");
    }
  };

  return (

    <Page theme={pageTheme.tool}>
      <Header title={`Customer System`} type="Computer Repair System" >
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
        href="/"
        variant="contained"  
        > 
        ออกจากระบบ 
        </Button>

      </Header>
      <Content>
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
                      value={customername}
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
                    checkresearch();
                    setSearch(true);

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
              {search ? (
                <div>
                  {  checkcustomername ? (
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
                          <TableCell align="center">เจ้าหน้าที่แจ้งซ่อม</TableCell>

                           
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {customer.filter((filter: any) => filter.customername.includes(customername)).map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.edges?.title?.titlename}</TableCell>
                              <TableCell align="center">{item.customername}</TableCell>
                              <TableCell align="center">{item.address}</TableCell>
                              <TableCell align="center">{item.phonenumber}</TableCell>
                              <TableCell align="center">{item.identificationnumber}</TableCell>
                              <TableCell align="center">{item.edges?.gender?.gendername}</TableCell>
                              <TableCell align="center">{item.edges?.personal.personalname}</TableCell>
                                                                  
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : customername == "" ? (
                      <div>
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
                              <TableCell align="center">เจ้าหน้าที่แจ้งซ่อม</TableCell>

                                
                              </TableRow>
                            </TableHead>
                            <TableBody>

                              {customer.map((item: any) => (
                                <TableRow key={item.id}>
                                 <TableCell align="center">{item.id}</TableCell>
                                <TableCell align="center">{item.edges?.title?.titlename}</TableCell>
                                <TableCell align="center">{item.customername}</TableCell>
                                <TableCell align="center">{item.address}</TableCell>
                                <TableCell align="center">{item.phonenumber}</TableCell>
                                <TableCell align="center">{item.identificationnumber}</TableCell>
                                <TableCell align="center">{item.edges?.gender?.gendername}</TableCell>
                                <TableCell align="center">{item.edges?.personal.personalname}</TableCell>
                                 
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