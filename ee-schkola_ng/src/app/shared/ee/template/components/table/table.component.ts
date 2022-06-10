import {Component, DoCheck, Input} from '@angular/core';
import {MatTableDataSource} from '@angular/material/table';
import {TableDataService} from '../../services/data.service';

@Component({
    selector: 'app-table',
    templateUrl: './table.component.html',
    styleUrls: ['./table.component.scss'],
    providers: [],
})

export class TableComponent implements DoCheck {
    @Input() displayedColumns: any[];

    dataSources: MatTableDataSource<any>;

    constructor(public tableDataService: TableDataService) { }

    ngDoCheck(): void {
        this.dataSources = new MatTableDataSource(this.tableDataService.retrieveItemFromCache());
    }
}
