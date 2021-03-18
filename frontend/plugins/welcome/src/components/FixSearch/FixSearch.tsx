//style
import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Page, pageTheme, Header, Content} from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
//api
import { DefaultApi } from '../../api/apis';
//entity
import { EntFix } from '../../api/models/EntFix';
//table
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import moment from 'moment'
import Paper from '@material-ui/core/Paper';
//alert
import Swal from 'sweetalert2'
//icon
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';

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


export default function FixSearch() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [fix, setFix] = useState<EntFix[]>([])
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);
  const [checkqueue, setCheckQueue] = useState(false);
  const [queue, setQueue] = useState(String);

  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบค้นหาข้อมูลบันทึกการแจ้งซ่อมสินค้า' };

  useEffect(() => {
    const getFixs = async () => {
      const res = await api.listFix({ offset: 0 });
      setLoading(false);
      setFix(res);
    };
    getFixs();
  }, [loading]);

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  const queuehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setCheckQueue(false);
    setQueue(event.target.value as string);
    if (event.target.value == "") {
      cleardata();
    }
  };

  const cleardata = async () => {
    const res = await api.listFix({ offset: 0 });
    setLoading(false);
    setFix(res);
  }
  
  const checkresearch = async () => {
    var check = false;
    fix.map(item => {
      if (queue != "") {
        if (item.queue?.includes(queue)) {
            setCheckQueue(true);
          alertMessage("success", "ค้นหาข้อมูลสำเร็จ");
          check = true;
          fix.splice(0, fix.length);
          fix.push(item);
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูลที่ค้นหา");
    }
    console.log(checkqueue)
    if (queue == "") {
      alertMessage("info", "กรุณากรอกเลขที่คิวเพื่อทำการค้นหา");
    }
  };

  return (

    <Page theme={pageTheme.tool}>
      <Header title="ระบบค้นหาบันทึกการแจ้งซ่อมคอมพิวเตอร์" type="ระบบแจ้งซ่อมคอมพิวเตอร์" > 
      <div>&nbsp;&nbsp;&nbsp;</div>
        <Button 
        style={{ marginLeft: 20 }} 
        href="/Tablefix"
        variant="contained"
        color="primary"
        startIcon={<HomeRoundedIcon/>}  
        > 
        กลับหน้าตารางข้อมูลการแจ้งซ่อม 
        </Button>

        <Button 
        style={{ marginLeft: 20 }} 
        href="/"
        variant="contained"  
        startIcon={<ExitToAppIcon/>}
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
                    ค้นหาข้อมูลการแจ้งซ่อมสินค้า
            </h1>
                </div>

                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}><strong>กรุณากรอกเลขที่คิวที่ต้องการค้นหา</strong></div>
                    <TextField
                      id="queue"
                      value={queue}
                      onChange={queuehandlehange}
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
                            <TableCell align="center">No</TableCell>
                            <TableCell align="center">แอดมิน</TableCell>
                            <TableCell align="center">ชื่อลูกค้า</TableCell>
                            <TableCell align="center">ประเภทคอมพิวเตอร์</TableCell>
                            <TableCell align="center">แบรนด์</TableCell>
                            <TableCell align="center">หมายเลขผลิตภัณฑ์</TableCell>
                            <TableCell align="center">วันที่รับแจ้งซ่อม</TableCell>
                            <TableCell align="center">รายละเอียดการแจ้งซ่อม/ปัญา</TableCell>
                            <TableCell align="center">ลำดับคิว</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                        {fix.map((item: any) => (
                            <TableRow key={item.id}>
                                <TableCell align="center">{item.id}</TableCell>
                                <TableCell align="center">{item.edges?.personal?.personalname}</TableCell>
                                <TableCell align="center">{item.edges?.customer?.customername}</TableCell>
                                <TableCell align="center">{item.edges?.fixcomtype?.fixcomtypename}</TableCell>
                                <TableCell align="center">{item.edges?.fixbrand?.fixbrandname}</TableCell>
                                <TableCell align="center">{item.productnumber}</TableCell>
                                <TableCell align="center">{moment(item.date).format("DD/MM/YYYY HH.mm น.")}</TableCell>
                                <TableCell align="center">{item.problemtype}</TableCell>
                                <TableCell align="center">{item.queue}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
              </Paper>
            </Grid>
         </Grid>
      </Content>
    </Page>
  );

}