import { TestBed } from '@angular/core/testing';
import { SearchComponent } from './search.component';
import { mount } from 'cypress/angular'
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

beforeEach(() => TestBed.configureTestingModule({
  imports: [HttpClientTestingModule], 
  providers: [SearchComponent]
}));

describe('SearchComponent', () => {
  it('mounts', () => {
    mount(SearchComponent)
  })
})