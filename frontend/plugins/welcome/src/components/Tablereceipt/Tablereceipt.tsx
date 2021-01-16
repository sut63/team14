//style
import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Content, ContentHeader, Header, Page, pageTheme } from '@backstage/core';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';

//table
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';

//api
import { DefaultApi } from '../../api/apis';

//Entity
import { EntReceipt } from '../../api';

//icon
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import DeleteIcon from '@material-ui/icons/Delete';
import NoteTwoToneIcon from '@material-ui/icons/NoteTwoTone';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  table: {
    minWidth: 650,
  },
  buttonRight: {
    marginLeft: theme.spacing(150),
    marginBottom: theme.spacing(2),
  },
  }),
);
 
export default function Tablepersonal() {
  const classes = useStyles();
  const http = new DefaultApi();

  const [loading, setLoading] = useState(true);

  const [receipts, setReceipts] = useState<EntReceipt[]>([]);
  
  useEffect(() => {
    const gettReceipts = async () => {
      const res:any = await http.listReceipt({ limit: 10, offset: 0 });
      setLoading(true);
      setReceipts(res);
      console.log(res);
    };
    gettReceipts();
  }, [loading]);
  
  // delete button
  const deleteReceipts = async (id: number) => {
    const res = await http.deleteReceipt({ id: id });
    setLoading(true);
  };

  // clear input form
  function clear() {
    setReceipts([]);
  }
  
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`ระบบออกใบเสร็จ`} type="ระบบแจ้งซ่อมคอมพิวเตอร์" >
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
      <ContentHeader title="ตารางแสดงผลข้อมูล">
      <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="secondary" href="/createreceipt" startIcon={<NoteTwoToneIcon />}> สร้างใบเสร็จ </Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/Group14" startIcon={<HomeRoundedIcon/>}> กลับหน้าหลัก </Button>
      </ContentHeader>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">รหัสลูกค้า</TableCell>
           <TableCell align="center">ชื่อ-นามสกุล</TableCell>
           <TableCell align="center">เบอร์โทรศัพท์</TableCell>
           <TableCell align="center">รายละเอียดการซ่อม</TableCell>
           <TableCell align="center">ราคา</TableCell>
           <TableCell align="center">ประเภทการจ่ายเงิน</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
       {receipts.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.customer?.id}</TableCell>
             <TableCell align="center">{item.edges?.customer?.customername}</TableCell>
             <TableCell align="center">{item.edges?.customer?.phonenumber}</TableCell>
             <TableCell align="center">{item.edges?.adminrepair?.equipmentdamate}</TableCell>
             <TableCell align="center">{item.edges?.product?.price}</TableCell>
             <TableCell align="center">{item.edges?.paymenttype?.typename}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                   deleteReceipts(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
  </Page>
);
}