import { DatePipe } from '@angular/common';
import { Pipe, PipeTransform } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import {TableDataService} from '@template/services/data.service';

@Pipe({
    name: 'DateTimeTranslation',
    pure: false
})
export class DateTimeTranslationPipe implements PipeTransform {

    constructor(private translateService: TranslateService, private dataService: TableDataService) {
        this.dataService.datePattern = 'shortDate';
    }

    transform(value: Date, pattern: string = this.dataService.datePattern): any {
        const datePipe: DatePipe = new DatePipe(
            this.translateService.currentLang ? this.translateService.currentLang : this.translateService.getDefaultLang());
        return datePipe.transform(value, pattern);
    }
}
