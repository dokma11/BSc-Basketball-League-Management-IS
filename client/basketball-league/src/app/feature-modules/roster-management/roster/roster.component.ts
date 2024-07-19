import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, OnInit, ViewChild } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { MatSelect } from '@angular/material/select';
import { MatSnackBar } from '@angular/material/snack-bar';
import { map, Observable, of, startWith, Subject, take, takeUntil } from 'rxjs';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';

@Component({
  selector: 'app-roster',
  templateUrl: './roster.component.html',
  styleUrls: ['./roster.component.css'],
  animations: [
      trigger("fadeIn", [
        transition(":enter", [
            style({ opacity: 0, transform: "translateX(-40px)" }),
            animate(
                "0.5s ease",
                style({ opacity: 1, transform: "translateX(0)" }),
            ),
        ]),
      ]),
      trigger('buttonState', [
        state('clicked', style({
          transform: 'scale(0.9)',
          opacity: 0.5
        })),
        transition('* => clicked', [
          animate('200ms')
        ]),
        transition('clicked => idle', [
          animate('200ms')
        ])
      ]),
  ],
})
export class RosterComponent implements OnInit{
  user: User | undefined;
  backgroundSize: string = '100% 100%';
  //requests: PersonalTourRequest[] = [];   ovde treba da budu zahtevi za trejdove bas
  tourRequestsButtonState: string = "";
  handledRequestsReportButtonState: string = "";
  requestsReportButtonState: string = "";

  private dialogRef: any;
  public teams: string[] = [
    'Brooklyn Nets',
    'Golden State Warriors',
    'Los Angeles Lakers',
    'Los Angeles Clippers',
    'New Orleans Pelicans',
    'New York Knicks',
    'Oklahoma City Thunder',
    'San Antonio Spurs',
    'Boston Celtics',
    'Denver Nuggets',
    'Minnesota Timberwolves',
    'Cleveland Cavaliers',
    'Philadelphia 76ers',
    'Phoenix Suns',
    'Sacramento Kings',
    'Indiana Pacers',
    'Dallas Mavericks',
    'Miami Heat',
    'Orlando Magic',
    'Chicago Bulls',
    'Atlanta Hawks',
    'Toronto Raptors',
    'Charlotte Hornets',
    'Washington Wizards',
    'Detroit Pistons',
    'Utah Jazz',
    'Houston Rockets',
    'Memphis Grizzlies',
    'Portland Trail Blazers',
    'Milwaukee Bucks'
  ];

  public teamCtrl: FormControl<string | null> = new FormControl<string | null>('');
  public teamFilterCtrl: FormControl<string | null> = new FormControl<string | null>('');

  public filteredTeams: Observable<string[]> = of(this.teams);

  @ViewChild('singleSelect', { static: true }) singleSelect: MatSelect | undefined;

  protected _onDestroy = new Subject<void>();

  constructor(private authService: AuthService,
              private dialog: MatDialog,
              private snackBar: MatSnackBar,) {

  }

  ngOnInit(): void {
    this.getRequests();

    this.teamCtrl.setValue('');

    this.filteredTeams = this.teamFilterCtrl.valueChanges.pipe(
      startWith(''),
      map(value => this.filterTeams(value))
    );
  }

  private filterTeams(value: any): string[] {
    const filterValue = value?.toLowerCase() || '';
    return this.teams.filter(team => team.toLowerCase().includes(filterValue));

    // TODO: Azurirati listu rosterom odgovarajuce ekipe
  }

  ngAfterViewInit() {
    this.setInitialValue();
  }

  ngOnDestroy() {
    this._onDestroy.next();
    this._onDestroy.complete();
  }

  protected setInitialValue() {
    this.filteredTeams
      .pipe(take(1), takeUntil(this._onDestroy))
      .subscribe(() => {
        this.singleSelect!.compareWith = (a: string, b: string) => a.toLowerCase() === b.toLowerCase();
      });
  }

  getRequests() {
    // this.authService.user$.subscribe(user => {
    //   this.user = user;
    //   if(this.user.role === 'GUEST'){
    //     this.toursService.getGuestsTourRequests(this.user.id).subscribe({
    //       next: (result: PersonalTourRequest[] | PersonalTourRequest) => {
    //         if(Array.isArray(result)){
    //           this.requests = result;
    //         }
    //       }
    //     });
    //   }
    //   else{
    //     this.toursService.getTourRequestsOnHold().subscribe({
    //       next: (result: PersonalTourRequest[] | PersonalTourRequest) => {
    //         if(Array.isArray(result)){
    //           this.requests = result;
    //         }
    //       }
    //     });
    //   }
    // });
  }

  handleDialogClosed(result: any) {
    this.getRequests();
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }
}
