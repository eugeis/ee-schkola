import {Attendance} from '@schkola/student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {AttendanceDataService} from '../../service/attendance-data.service';

@Component({
  selector: 'app-attendance-view',
  templateUrl: './attendance-entity-view.component.html',
  styleUrls: ['./attendance-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: AttendanceDataService}]
})

export class AttendanceViewComponent implements OnInit {

    attendance: Attendance;

    constructor(public attendanceDataService: AttendanceDataService) {}

    ngOnInit(): void {
        this.attendance = this.attendanceDataService.getFirst();
        this.attendanceDataService.checkRoute(this.attendance);
    }
}




