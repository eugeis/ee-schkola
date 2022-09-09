import {Attendance, AttendanceState} from '../student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {AttendanceDataService} from '../../services/attendance-data.service';

@Component({
  selector: 'app-attendance-view',
  templateUrl: './attendance-view.component.html',
  styleUrls: ['./attendance-view.component.scss'],
  providers: [{provide: TableDataService, useClass: AttendanceDataService}]
})

export class AttendanceViewComponent implements OnInit {

    attendancestateEnum = this.attendanceDataService.loadEnumElement(AttendanceState);
    attendance: Attendance;

    constructor(public attendanceDataService: AttendanceDataService) {}

    ngOnInit(): void {
        this.attendance = this.attendanceDataService.getFirst();
        this.attendanceDataService.checkRoute(this.attendance);
    }
}




