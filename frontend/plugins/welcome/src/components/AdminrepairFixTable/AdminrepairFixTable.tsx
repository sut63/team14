import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import { EntFix } from '../../api/models/EntFix';
import moment from 'moment'
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [loading, setLoading] = useState(true);
 const [fixs, setFixs] = useState<EntFix[]>([]);
 
 useEffect(() => {
    const getfixs = async () => {
      const res:any = await api.listFix({ limit: 10, offset: 0 });
      setLoading(true);
      setFixs(res);
      console.log(res);
    };
    getfixs();
  }, [loading]);
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
       <TableRow>
           <TableCell align="center">No</TableCell>
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
         {fixs.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.customer?.customername}</TableCell>
             <TableCell align="center">{item.edges?.fixcomtype?.fixcomtypename}</TableCell>
             <TableCell align="center">{item.edges?.fixbrand?.fixbrandname}</TableCell>
             <TableCell align="center">{item.productnumber}</TableCell>
             <TableCell align="center">{moment(item.date).format("DD/MM/YYYY HH.mm น.")}</TableCell>
             <TableCell align="center">{item.problemtype}</TableCell>
             <TableCell align="center">{item.queue}</TableCell>
             <TableCell align="center">
             </TableCell>

           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
