import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LeagueNewsComponent } from './league-news.component';

describe('LeagueNewsComponent', () => {
  let component: LeagueNewsComponent;
  let fixture: ComponentFixture<LeagueNewsComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [LeagueNewsComponent]
    });
    fixture = TestBed.createComponent(LeagueNewsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
