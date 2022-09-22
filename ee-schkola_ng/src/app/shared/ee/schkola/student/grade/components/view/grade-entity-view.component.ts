import {Grade} from '@schkola/student/StudentApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {GradeDataService} from '../../service/grade-data.service';

@Component({
  selector: 'app-grade-view',
  templateUrl: './grade-entity-view.component.html',
  styleUrls: ['./grade-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: GradeDataService}]
})

export class GradeViewComponent implements OnInit {


    grade: Grade;

    constructor(@Inject(GradeDataService) public gradeDataService: GradeDataService) {}

    ngOnInit(): void {
        this.grade = this.gradeDataService.getFirst();
        
        this.gradeDataService.checkRoute(this.grade);
    }
}




