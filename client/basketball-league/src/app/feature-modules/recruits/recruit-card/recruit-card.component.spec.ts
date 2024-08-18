import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RecruitCardComponent } from './recruit-card.component';

describe('RecruitCardComponent', () => {
  let component: RecruitCardComponent;
  let fixture: ComponentFixture<RecruitCardComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [RecruitCardComponent]
    });
    fixture = TestBed.createComponent(RecruitCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
