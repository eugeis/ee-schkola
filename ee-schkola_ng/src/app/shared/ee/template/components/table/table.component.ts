import {Component, Input, OnInit} from '@angular/core';
import {MatTableDataSource} from '@angular/material/table';
import {TableDataService} from '../../services/data.service';

@Component({
    selector: 'app-table',
    templateUrl: './table.component.html',
    styleUrls: ['./table.component.scss'],
})

export class TableComponent implements OnInit {
    @Input() displayedColumns: any[];

    dataSources: MatTableDataSource<any>;

    constructor(public tableDataService: TableDataService) { }

    ngOnInit() {
        this.dataSources = new MatTableDataSource(this.tableDataService.retrieveItemFromCache());
    }
}
