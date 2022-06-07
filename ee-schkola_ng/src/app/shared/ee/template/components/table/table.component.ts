import {Component, DoCheck, Input} from '@angular/core';
import {MatTableDataSource} from '@angular/material/table';
import {ELEMENT_DATA} from '../../services/template.service';

@Component({
    selector: 'app-table',
    templateUrl: './table.component.html',
    styleUrls: ['./table.component.scss'],
})

export class TableComponent implements DoCheck {
    @Input() displayedColumns: any[];

    dataSources: MatTableDataSource<any>;

    constructor() { }

    ngDoCheck(): void {
        this.dataSources = new MatTableDataSource(ELEMENT_DATA);
    }
}
