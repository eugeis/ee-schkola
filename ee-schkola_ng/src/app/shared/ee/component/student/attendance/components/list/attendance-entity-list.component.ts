import {Attendance} from '../student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {AttendanceDataService} from '../../services/attendance-data.service';

@Component({
  selector: 'app-attendance-list',
  templateUrl: './attendance-view.component.html',
  styleUrls: ['./attendance-view.component.scss'],
  providers: [{provide: TableDataService, useClass: AttendanceDataService}]
})

export class AttendanceViewComponent implements OnInit {

    attendance: Attendance = new Attendance();

    tableHeader: Array<String> = [];

    constructor(public attendanceDataService: AttendanceDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'student', 'date', 'course', 'hours', 'state', 'token', 'id'];
    }
}




