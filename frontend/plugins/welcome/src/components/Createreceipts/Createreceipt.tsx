//style
import React, {  useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField, Button, InputAdornment } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Alert, AlertTitle } from '@material-ui/lab';

//api
import { DefaultApi } from '../../api/apis';

//Entity
import { EntPersonal, EntPaymentType, EntReceipt, EntAdminrepair, EntCustomer, EntProduct } from '../../api';

//icon
import NoteTwoToneIcon from '@material-ui/icons/NoteTwoTone';
import AddCircleTwoToneIcon from '@material-ui/icons/AddCircleTwoTone';
import CancelTwoToneIcon from '@material-ui/icons/CancelTwoTone';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';

//cookie
import { Cookies } from '../WelcomePage/Cookie';

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
   marginLeft: '40px',
  },
  textField: {
   width: 500 ,
   marginLeft:7,
   marginRight:-7,
  },
  select: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
  },
  paper: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(1),
  },
  pa: {
    marginTop: theme.spacing(2),
  },
  }),
);

var ck = new Cookies()
var cookieName = ck.GetCookie()
var cookieID = ck.GetID()

export default function Personalpage() {
  const classes = useStyles();
  const http = new DefaultApi();

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [personals, setPersonals] = React.useState<EntPersonal[]>([]);
  const [adminrepairs, setAdminrepairs] = React.useState<EntAdminrepair[]>([]);
  const [paymenttypes, setPaymenttypes] = React.useState<EntPaymentType[]>([]);
  const [customers, setCustomers] = React.useState<EntCustomer[]>([]);
  const [products, setProducts] = React.useState<EntProduct[]>([]);

  const [personal, setPersonal] = useState(Number);
  const [adminrepair, setAdminrepair] = useState(Number);
  const [paymenttype, setPaymenttype] = useState(Number);
  const [customer, setCustomer] = useState(Number);
  const [product, setProduct] = useState(Number);

  const [addedTime, setAddedTime] = useState(String);
  

  useEffect(() => {
    const getCustomers = async () => {
      const res = await http.listCustomer({ limit: 10, offset: 0 });
      setLoading(false);
      setCustomers(res);
      console.log(res);
    };
    getCustomers();

    const getAdminrepairs = async () => {
      const res = await http.listAdminrepair({ limit: 10, offset: 0 });
      setLoading(false);
      setAdminrepairs(res);
      console.log(res);
    };
    getAdminrepairs();

    const getPaymenttypes = async () => {
      const res = await http.listPaymenttype({ limit: 10, offset: 0 });
      setLoading(false);
      setPaymenttypes(res);
      console.log(res);
    };
    getPaymenttypes();

    const getPersonals = async () => {
      const res = await http.listPersonal({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonals(res);
      console.log(res);
    };
    getPersonals();

  const getProducts = async () => {
    const res = await http.listProduct({ limit: 10, offset: 0 });
    setLoading(false);
    setProducts(res);
    console.log(res);
  };
  getProducts();

}, [loading]);

  const addedTimeChange = (event: any) => {
    setAddedTime(event.target.value as string);
  };
  
  const personalChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonal(event.target.value as number);
  };

  const paymenttypechange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPaymenttype(event.target.value as number);
  };

  const adminrepairchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAdminrepair(event.target.value as number);
  };

  const customerchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setCustomer(event.target.value as number);
  };

  const productchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProduct(event.target.value as number);
  };

  const CreateReceipt = async () => {
    if ( (customer != null) && (personal != null) && (adminrepair != null) && (paymenttype != null) && (product != null) && (addedTime != "") &&(addedTime != null)){
    const receipt = {
      customer: customer,
      personal: Number(cookieID),
      adminrepair: adminrepair,
      paymentType: paymenttype,
      product: product,
      added : addedTime + ":00+07:00", 
    };
    console.log(receipt)
    const res: any = await http.createReceipt({ receipt: receipt });
      setStatus(true);
      if (res.id != '') {
        setAlert(true);
      } 
    }
      else {
        setStatus(true);
        setAlert(false);
      }
  };

return (
  <Page theme={pageTheme.tool}>
    <Header
      title="ระบบออกใบเสร็จ" type="ระบบแจ้งซ่อมคอมพิวเตอร์">
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
      <ContentHeader title="สร้างใบเสร็จ">
      <div>&nbsp;&nbsp;&nbsp;</div>
        <Button  
        onClick={() => {CreateReceipt();}}
        variant="contained" 
        color="secondary" 
        startIcon={<NoteTwoToneIcon/>}
        > 
        บันทึกข้อมูล 
        </Button>
      <div>&nbsp;&nbsp;&nbsp;</div>
        <Button 
        style={{ marginLeft: 20 }} 
        component={RouterLink} 
        to="/Tablereceipt" 
        variant="contained"
        color="primary"
        startIcon={<CancelTwoToneIcon/>}
        > 
        ย้อนกลับ 
        </Button>
      </ContentHeader>
      {status ? ( 
        <div>
          {alert ? ( 
              <Alert severity="success"  onClose={() => { }}> 
                <AlertTitle> บันทึกข้อมูลสำเร็จ </AlertTitle></Alert>) 
        : (     
          <Alert severity="error" onClose={() => { setStatus(false) }}> 
            <AlertTitle> ไม่สามารถบันทึกข้อมูลได้ กรุณาลองใหม่อีกครั้ง </AlertTitle></Alert>)}
        </div>
          ) : null}
      <div className={classes.root}>
        <form noValidate autoComplete="off">
          <FormControl
            fullWidth
            className={classes.margin}
            variant="outlined"
            size="small"
          >
            <div className={classes.paper}><strong>รหัสลูกค้า</strong></div>
              <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              labelId="nametitle-label"
              id="customerid"
              value={customer}
              onChange={customerchange}
            >
              {customers.map((item: EntCustomer) => (
                <MenuItem value={item.id}>{item.id}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>ชื่อ-นามสกุล</strong></div>
              <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              labelId="nametitle-label"
              id="namecustomer"
              value={customer}
              onChange={customerchange}
            >
              {customers.map((item: EntCustomer) => (
                <MenuItem value={item.id}>{item.customername}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>เบอร์โทรศัพท์</strong></div>
              <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              labelId="nametitle-label"
              id="phonenumber"
              value={customer}
              onChange={customerchange}
            >
              {customers.map((item: EntCustomer) => (
                <MenuItem value={item.id}>{item.phonenumber}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>รายละเอียดการซ่อม</strong></div>
            <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              id="adminrepair"
              value={adminrepair}
              onChange={adminrepairchange}
            >
              {adminrepairs.map((item: EntAdminrepair) => (
                <MenuItem value={item.id}>{item.equipmentdamate}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>ราคา</strong></div>
            <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              id="price"
              value={product}
              onChange={productchange}
            >
              {products.map((item: EntProduct) => (
                <MenuItem value={item.id}>{item.price}</MenuItem>
              ))}
            </Select>
            

            <div className={classes.paper}><strong>ประเภทการจ่ายเงิน</strong></div>
            <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              id="paymenttype"
              value={paymenttype}
              onChange={paymenttypechange}
            >
              <InputLabel className={classes.insideLabel}>เลือกประเภทการจ่ายเงิน</InputLabel>

              {paymenttypes.map((item: EntPaymentType) => (
                <MenuItem value={item.id}>{item.typename}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>วันที่รับแจ้งซ่อม</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <AddCircleTwoToneIcon />
                </InputAdornment>
              ),
            }}
              id="date"
              variant="standard"
              color="secondary"
              type="date"
              size="medium"
              value={addedTime}
              onChange={addedTimeChange}
            />
            
          </FormControl>
        </form>
      </div>
    </Content>
   </Page>
 );
}